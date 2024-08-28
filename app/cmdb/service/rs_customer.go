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

type RsCustomer struct {
	service.Service
}

// GetPage 获取RsCustomer列表
func (e *RsCustomer) GetPage(c *dto.RsCustomerGetPageReq, p *actions.DataPermission, list *[]models.RsCustomer, count *int64) error {
	var err error
	var data models.RsCustomer

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsCustomerService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsCustomer对象
func (e *RsCustomer) Get(d *dto.RsCustomerGetReq, p *actions.DataPermission, model *models.RsCustomer) error {
	var data models.RsCustomer

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsCustomer error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RsCustomer对象
func (e *RsCustomer) Insert(c *dto.RsCustomerInsertReq) (id int, err error) {

	var data models.RsCustomer
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsCustomerService Insert error:%s \r\n", err)
		return 0, err
	}
	return data.Id, nil
}

// Update 修改RsCustomer对象
func (e *RsCustomer) Update(c *dto.RsCustomerUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RsCustomer{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RsCustomerService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RsCustomer
func (e *RsCustomer) Remove(d *dto.RsCustomerDeleteReq, p *actions.DataPermission) error {
	var data models.RsCustomer

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRsCustomer error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
