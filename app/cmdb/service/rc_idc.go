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

type RcIdc struct {
	service.Service
}

// GetPage 获取RcIdc列表
func (e *RcIdc) GetPage(c *dto.RcIdcGetPageReq, p *actions.DataPermission, list *[]models.RcIdc, count *int64) error {
	var err error
	var data models.RcIdc

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RcIdcService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RcIdc对象
func (e *RcIdc) Get(d *dto.RcIdcGetReq, p *actions.DataPermission, model *models.RcIdc) error {
	var data models.RcIdc

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRcIdc error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RcIdc对象
func (e *RcIdc) Insert(c *dto.RcIdcInsertReq) error {
    var err error
    var data models.RcIdc
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RcIdcService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RcIdc对象
func (e *RcIdc) Update(c *dto.RcIdcUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.RcIdc{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("RcIdcService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除RcIdc
func (e *RcIdc) Remove(d *dto.RcIdcDeleteReq, p *actions.DataPermission) error {
	var data models.RcIdc

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveRcIdc error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
