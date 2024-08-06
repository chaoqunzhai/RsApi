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

type RsContacts struct {
	service.Service
}

// GetPage 获取RsContacts列表
func (e *RsContacts) GetPage(c *dto.RsContactsGetPageReq, p *actions.DataPermission, list *[]models.RsContacts, count *int64) error {
	var err error
	var data models.RsContacts

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsContactsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsContacts对象
func (e *RsContacts) Get(d *dto.RsContactsGetReq, p *actions.DataPermission, model *models.RsContacts) error {
	var data models.RsContacts

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsContacts error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RsContacts对象
func (e *RsContacts) Insert(c *dto.RsContactsInsertReq) error {
    var err error
    var data models.RsContacts
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsContactsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RsContacts对象
func (e *RsContacts) Update(c *dto.RsContactsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.RsContacts{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("RsContactsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除RsContacts
func (e *RsContacts) Remove(d *dto.RsContactsDeleteReq, p *actions.DataPermission) error {
	var data models.RsContacts

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveRsContacts error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
