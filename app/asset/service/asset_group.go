package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/asset/models"
	"go-admin/app/asset/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type AssetGroup struct {
	service.Service
}

// GetPage 获取AssetGroup列表
func (e *AssetGroup) GetPage(c *dto.AssetGroupGetPageReq, p *actions.DataPermission, list *[]models.AssetGroup, count *int64) error {
	var err error
	var data models.AssetGroup

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AssetGroupService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AssetGroup对象
func (e *AssetGroup) Get(d *dto.AssetGroupGetReq, p *actions.DataPermission, model *models.AssetGroup) error {
	var data models.AssetGroup

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAssetGroup error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AssetGroup对象
func (e *AssetGroup) Insert(c *dto.AssetGroupInsertReq) error {
	var err error
	var data models.AssetGroup
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AssetGroupService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AssetGroup对象
func (e *AssetGroup) Update(c *dto.AssetGroupUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AssetGroup{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AssetGroupService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AssetGroup
func (e *AssetGroup) Remove(d *dto.AssetGroupDeleteReq, p *actions.DataPermission) error {
	var data models.AssetGroup

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAssetGroup error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
