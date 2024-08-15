package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/common/utils"
	"gorm.io/gorm"

	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type RsContract struct {
	service.Service
}

// GetPage 获取RsContract列表
func (e *RsContract) GetPage(c *dto.RsContractGetPageReq, p *actions.DataPermission, list *[]models.RsContract, count *int64) error {
	var err error
	var data models.RsContract

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsContractService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsContract对象
func (e *RsContract) Get(d *dto.RsContractGetReq, p *actions.DataPermission, model *models.RsContract) error {
	var data models.RsContract

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsContract error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	return nil
}

// Insert 创建RsContract对象
func (e *RsContract) Insert(c *dto.RsContractInsertReq) error {
	var err error
	var data models.RsContract
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsContractService Insert error:%s \r\n", err)
		return err
	}
	for _, row := range c.BandwidthFees {
		var bandRow models.RsBandwidthFees
		row.Generate(&bandRow)
		bandRow.ContractId = data.Id
		err = e.Orm.Create(&bandRow).Error
	}
	return nil
}

// Update 修改RsContract对象
func (e *RsContract) Update(c *dto.RsContractUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RsContract{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RsContractService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	hasIds := make([]int, 0)
	var RsBandwidthFeesList []models.RsBandwidthFees
	e.Orm.Model(&models.RsBandwidthFees{}).Select("id").Where("contract_id = ?", c.Id).Find(&RsBandwidthFeesList)

	for _, row := range RsBandwidthFeesList {
		hasIds = append(hasIds, row.Id)
	}

	updateId := make([]int, 0)
	for _, row := range c.BandwidthFees {
		var bandRow models.RsBandwidthFees
		if row.Id > 0 { //更新
			e.Orm.Model(&bandRow).First(&bandRow, row.GetId())
			row.Generate(&bandRow)
			bandRow.ContractId = data.Id
			updateId = append(updateId, row.Id)
			e.Orm.Save(&bandRow)
		} else { //创建
			row.Generate(&bandRow)
			bandRow.ContractId = data.Id
			err = e.Orm.Create(&bandRow).Error
		}
	}
	diffIds := utils.DifferenceInt(hasIds, updateId)

	e.Orm.Model(&models.RsBandwidthFees{}).Where("id in ?", diffIds).Delete(&models.RsBandwidthFees{})
	return nil
}

// Remove 删除RsContract
func (e *RsContract) Remove(d *dto.RsContractDeleteReq, p *actions.DataPermission) error {
	var data models.RsContract

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRsContract error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	//删除对应的带宽配置信息
	var bandWih models.RsBandwidthFees
	e.Orm.Model(&bandWih).Delete(&bandWih, d.GetId())
	return nil
}
