package service

import (
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/asset/models"
	"go-admin/app/asset/service/dto"
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/utils"
	"gorm.io/gorm"
	"strings"
)

type AssetOutboundOrder struct {
	service.Service
}

// GetPage 获取AssetOutboundOrder列表
func (e *AssetOutboundOrder) GetPage(c *dto.AssetOutboundOrderGetPageReq, p *actions.DataPermission, list *[]models.AssetOutboundOrder, count *int64) error {
	var err error
	var data models.AssetOutboundOrder

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AssetOutboundOrderService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AssetOutboundOrder对象
func (e *AssetOutboundOrder) Get(d *dto.AssetOutboundOrderGetReq, p *actions.DataPermission, model *models.AssetOutboundOrder) error {
	var data models.AssetOutboundOrder

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAssetOutboundOrder error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AssetOutboundOrder对象
func (e *AssetOutboundOrder) Insert(c *dto.AssetOutboundOrderInsertReq) error {
	var err error
	var data models.AssetOutboundOrder
	var userModel models2.SysUser
	e.Orm.Model(&models2.SysUser{}).Where("user_id = ?", c.CreateBy).Limit(1).Find(&userModel)

	c.Generate(&data)
	if len(c.Combination) > 0 {
		CombinationIds := utils.RemoveRepeatInt(c.Combination)

		data.CombinationId = strings.Join(utils.IntToStringArray(CombinationIds), ",")
	}
	err = e.Orm.Create(&data).Error

	Code := fmt.Sprintf("CK%08d", data.Id)
	if err != nil {
		e.Log.Errorf("AssetOutboundOrderService Insert error:%s \r\n", err)
		return err
	}

	e.Orm.Model(&models.AssetOutboundOrder{}).Where("id = ?", data.Id).Updates(map[string]interface{}{
		"code": Code,
	})
	//查询关联的值, 以下资产必须是在库的
	//查询到combination 的组合
	recordingIds := make([]int, 0)

	if len(c.Combination) > 0 {
		CombinationIds := utils.RemoveRepeatInt(c.Combination)

		e.Orm.Model(&models.Combination{}).Where("id in ?", CombinationIds).Updates(map[string]interface{}{
			"status": 2,
		})

		e.Orm.Model(&models.AdditionsWarehousing{}).Where("combination_id in ?", CombinationIds).Updates(map[string]interface{}{
			"out_id": data.Id,
			"status": 2,
		})
		for _, i := range CombinationIds {
			e.Orm.Create(&models.AssetRecording{
				User:      userModel.Username,
				Type:      2,
				AssetType: 2,
				BindOrder: Code,
				Info:      "组合出库",
				AssetId:   i,
			})
			var CombinationModel models.Combination
			e.Orm.Model(&models.Combination{}).Where("id = ?", i).First(&CombinationModel)
			if CombinationModel.Id > 0 {
				//把对应的资产改为待上架
				e.Orm.Model(&models2.Host{}).Where("id = ?", CombinationModel.HostId).Updates(map[string]interface{}{
					"status": 3,
				})
			}
		}

	}
	//查询到asset中的资产
	if len(c.Asset) > 0 {
		var assetList []models.AdditionsWarehousing
		e.Orm.Model(&models.AdditionsWarehousing{}).Where("id in ? and status = 1", c.Asset).Updates(map[string]interface{}{
			"out_id": data.Id,
			"status": 2,
		}).Find(&assetList)
		for _, i := range assetList {
			recordingIds = append(recordingIds, i.Id)
		}
	}

	//进行操作日志记录
	recordingIds = utils.RemoveRepeatInt(recordingIds)

	for _, i := range recordingIds {
		e.Orm.Create(&models.AssetRecording{
			User:      userModel.Username,
			Type:      2,
			AssetType: 1,
			BindOrder: Code,
			AssetId:   i,
			CreateBy:  data.CreateBy,
		})
	}
	return nil
}

// Update 修改AssetOutboundOrder对象
func (e *AssetOutboundOrder) Update(c *dto.AssetOutboundOrderUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AssetOutboundOrder{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AssetOutboundOrderService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	return nil
}
