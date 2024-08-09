package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/asset/models"
	"go-admin/app/asset/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type AssetInbound struct {
	service.Service
}

// GetPage 获取AssetInbound列表
func (e *AssetInbound) GetPage(c *dto.AssetInboundGetPageReq, p *actions.DataPermission, list *[]models.AssetInbound, count *int64) error {
	var err error
	var data models.AssetInbound

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AssetInboundService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AssetInbound对象
func (e *AssetInbound) Get(d *dto.AssetInboundGetReq, p *actions.DataPermission, model *models.AssetInbound) error {
	var data models.AssetInbound

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAssetInbound error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AssetInbound对象
func (e *AssetInbound) Insert(c *dto.AssetInboundInsertReq) error {
	var err error
	var data models.AssetInbound
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AssetInboundService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AssetInbound对象
func (e *AssetInbound) Update(c *dto.AssetInboundUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AssetInbound{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AssetInboundService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AssetInbound
func (e *AssetInbound) Remove(d *dto.AssetInboundDeleteReq, p *actions.DataPermission) error {
	var data models.AssetInbound

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAssetInbound error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
