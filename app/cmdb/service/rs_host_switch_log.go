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

type RsHostSwitchLog struct {
	service.Service
}

// GetPage 获取RsHostSwitchLog列表
func (e *RsHostSwitchLog) GetPage(c *dto.RsHostSwitchLogGetPageReq, p *actions.DataPermission, list *[]models.RsHostSwitchLog, count *int64) error {
	var err error
	var data models.RsHostSwitchLog

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsHostSwitchLogService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsHostSwitchLog对象
func (e *RsHostSwitchLog) Get(d *dto.RsHostSwitchLogGetReq, p *actions.DataPermission, model *models.RsHostSwitchLog) error {
	var data models.RsHostSwitchLog

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsHostSwitchLog error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RsHostSwitchLog对象
func (e *RsHostSwitchLog) Insert(c *dto.RsHostSwitchLogInsertReq) error {
    var err error
    var data models.RsHostSwitchLog
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsHostSwitchLogService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RsHostSwitchLog对象
func (e *RsHostSwitchLog) Update(c *dto.RsHostSwitchLogUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.RsHostSwitchLog{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("RsHostSwitchLogService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除RsHostSwitchLog
func (e *RsHostSwitchLog) Remove(d *dto.RsHostSwitchLogDeleteReq, p *actions.DataPermission) error {
	var data models.RsHostSwitchLog

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveRsHostSwitchLog error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
