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

type RsHost struct {
	service.Service
}

// GetPage 获取RsHost列表
func (e *RsHost) GetPage(c *dto.RsHostGetPageReq, p *actions.DataPermission, list *[]models.RsHost, count *int64) error {
	var err error
	var data models.RsHost

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsHostService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsHost对象
func (e *RsHost) Get(d *dto.RsHostGetReq, p *actions.DataPermission, model *models.RsHost) error {
	var data models.RsHost

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsHost error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RsHost对象
func (e *RsHost) Insert(c *dto.RsHostInsertReq) error {
	var err error
	var data models.RsHost
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsHostService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RsHost对象
func (e *RsHost) Update(c *dto.RsHostUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RsHost{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RsHostService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RsHost
func (e *RsHost) Remove(d *dto.RsHostDeleteReq, p *actions.DataPermission) error {
	var data models.RsHost

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRsHost error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
