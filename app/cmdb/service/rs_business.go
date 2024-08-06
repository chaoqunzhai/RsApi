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

type RsBusiness struct {
	service.Service
}

// GetPage 获取RsBusiness列表
func (e *RsBusiness) GetPage(c *dto.RsBusinessGetPageReq, p *actions.DataPermission, list *[]models.RsBusiness, count *int64) error {
	var err error
	var data models.RsBusiness

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsBusinessService GetPage error:%s \r\n", err)
		return err
	}

	return nil
}

// Get 获取RsBusiness对象
func (e *RsBusiness) Get(d *dto.RsBusinessGetReq, p *actions.DataPermission, model *models.RsBusiness) error {
	var data models.RsBusiness

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsBusiness error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RsBusiness对象
func (e *RsBusiness) Insert(c *dto.RsBusinessInsertReq) error {
	var err error
	var data models.RsBusiness
	c.Generate(&data)
	data.Layer = 1
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsBusinessService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RsBusiness对象
func (e *RsBusiness) Update(c *dto.RsBusinessUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RsBusiness{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RsBusinessService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RsBusiness
func (e *RsBusiness) Remove(d *dto.RsBusinessDeleteReq, p *actions.DataPermission) error {
	var data models.RsBusiness

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRsBusiness error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
