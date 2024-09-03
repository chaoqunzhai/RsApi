package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type RsHostChargingDay struct {
	service.Service
}

// GetPage 获取RsHostChargingDay列表
func (e *RsHostChargingDay) GetPage(c *dto.RsHostChargingDayGetPageReq, p *actions.DataPermission, list *[]models.RsHostChargingDay, count *int64) error {
	var err error
	var data models.RsHostChargingDay

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsHostChargingDayService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsHostChargingDay对象
func (e *RsHostChargingDay) Get(d *dto.RsHostChargingDayGetReq, p *actions.DataPermission, model *models.RsHostChargingDay) error {
	var data models.RsHostChargingDay

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsHostChargingDay error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RsHostChargingDay对象
func (e *RsHostChargingDay) Insert(c *dto.RsHostChargingDayInsertReq) error {
	var err error
	var data models.RsHostChargingDay
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsHostChargingDayService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RsHostChargingDay对象
func (e *RsHostChargingDay) Update(c *dto.RsHostChargingDayUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RsHostChargingDay{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RsHostChargingDayService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RsHostChargingDay
func (e *RsHostChargingDay) Remove(d *dto.RsHostChargingDayDeleteReq, p *actions.DataPermission) error {
	var data models.RsHostChargingDay

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRsHostChargingDay error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
