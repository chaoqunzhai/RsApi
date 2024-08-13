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

type RsCustom struct {
	service.Service
}

// GetPage 获取RsCustom列表
func (e *RsCustom) GetPage(c *dto.RsCustomGetPageReq, p *actions.DataPermission, list *[]models.RsCustom, count *int64) error {
	var err error
	var data models.RsCustom

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsCustomService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsCustom对象
func (e *RsCustom) Get(d *dto.RsCustomGetReq, p *actions.DataPermission, model *models.RsCustom) error {
	var data models.RsCustom

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsCustom error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RsCustom对象
func (e *RsCustom) Insert(c *dto.RsCustomInsertReq) error {
	var err error
	var data models.RsCustom
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsCustomService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RsCustom对象
func (e *RsCustom) Update(c *dto.RsCustomUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RsCustom{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RsCustomService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RsCustom
func (e *RsCustom) Remove(d *dto.RsCustomDeleteReq, p *actions.DataPermission) error {
	var data models.RsCustom

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRsCustom error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
