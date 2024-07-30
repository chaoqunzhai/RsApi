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

type RsTag struct {
	service.Service
}

// GetPage 获取RsTag列表
func (e *RsTag) GetPage(c *dto.RsTagGetPageReq, p *actions.DataPermission, list *[]models.RsTag, count *int64) error {
	var err error
	var data models.RsTag

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsTagService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsTag对象
func (e *RsTag) Get(d *dto.RsTagGetReq, p *actions.DataPermission, model *models.RsTag) error {
	var data models.RsTag

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsTag error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RsTag对象
func (e *RsTag) Insert(c *dto.RsTagInsertReq) error {
    var err error
    var data models.RsTag
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsTagService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RsTag对象
func (e *RsTag) Update(c *dto.RsTagUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.RsTag{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("RsTagService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除RsTag
func (e *RsTag) Remove(d *dto.RsTagDeleteReq, p *actions.DataPermission) error {
	var data models.RsTag

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveRsTag error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
