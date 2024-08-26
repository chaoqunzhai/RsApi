package service

import (
	"errors"
	"fmt"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type RsIdc struct {
	service.Service
}

// GetPage 获取RsIdc列表
func (e *RsIdc) GetPage(c *dto.RsIdcGetPageReq, p *actions.DataPermission, list *[]models.RsIdc, count *int64) error {
	var err error
	var data models.RsIdc

	orm := e.Orm.Model(&data)
	if c.Search != "" {
		likeQ := fmt.Sprintf("number like '%%%s%%' or name like '%%%s%%' ", c.Search, c.Search)
		orm = orm.Where(likeQ)
	}
	err = orm.Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		actions.Permission(data.TableName(), p),
	).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsIdcService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsIdc对象
func (e *RsIdc) Get(d *dto.RsIdcGetReq, p *actions.DataPermission, model *models.RsIdc) error {
	var data models.RsIdc

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsIdc error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RsIdc对象
func (e *RsIdc) Insert(c *dto.RsIdcInsertReq) error {
	var err error
	var data models.RsIdc
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsIdcService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RsIdc对象
func (e *RsIdc) Update(c *dto.RsIdcUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RsIdc{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RsIdcService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RsIdc
func (e *RsIdc) Remove(d *dto.RsIdcDeleteReq, p *actions.DataPermission) error {
	var data models.RsIdc

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRsIdc error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
