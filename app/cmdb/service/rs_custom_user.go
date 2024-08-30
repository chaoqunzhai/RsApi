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

type RsCustomUser struct {
	service.Service
}

// GetPage 获取RsCustomUser列表
func (e *RsCustomUser) GetPage(c *dto.RsCustomUserGetPageReq, p *actions.DataPermission, list *[]models.RsCustomUser, count *int64) error {
	var err error
	var data models.RsCustomUser

	orm := e.Orm.Model(&data)
	if c.Search != "" {
		orm = orm.Where("user_name LIKE ? OR phone LIKE ? OR email LIKE ? ",
			fmt.Sprintf("%%%s%%", c.Search), fmt.Sprintf("%%%s%%", c.Search), fmt.Sprintf("%%%s%%", c.Search))
	}
	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsCustomUserService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsCustomUser对象
func (e *RsCustomUser) Get(d *dto.RsCustomUserGetReq, p *actions.DataPermission, model *models.RsCustomUser) error {
	var data models.RsCustomUser

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsCustomUser error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RsCustomUser对象
func (e *RsCustomUser) Insert(c *dto.RsCustomUserInsertReq) (modelId int, err error) {

	var data models.RsCustomUser
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsCustomUserService Insert error:%s \r\n", err)
		return 0, err
	}
	return modelId, nil
}

// Update 修改RsCustomUser对象
func (e *RsCustomUser) Update(c *dto.RsCustomUserUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RsCustomUser{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RsCustomUserService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RsCustomUser
func (e *RsCustomUser) Remove(d *dto.RsCustomUserDeleteReq, p *actions.DataPermission) error {
	var data models.RsCustomUser

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRsCustomUser error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
