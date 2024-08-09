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

type AssetCategory struct {
	service.Service
}

// GetPage 获取AssetCategory列表
func (e *AssetCategory) GetPage(c *dto.AssetCategoryGetPageReq, p *actions.DataPermission, list *[]models.AssetCategory, count *int64) error {
	var err error
	var data models.AssetCategory

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AssetCategoryService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AssetCategory对象
func (e *AssetCategory) Get(d *dto.AssetCategoryGetReq, p *actions.DataPermission, model *models.AssetCategory) error {
	var data models.AssetCategory

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAssetCategory error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AssetCategory对象
func (e *AssetCategory) Insert(c *dto.AssetCategoryInsertReq) error {
	var err error
	var data models.AssetCategory
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AssetCategoryService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AssetCategory对象
func (e *AssetCategory) Update(c *dto.AssetCategoryUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AssetCategory{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AssetCategoryService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AssetCategory
func (e *AssetCategory) Remove(d *dto.AssetCategoryDeleteReq, p *actions.DataPermission) error {
	var data models.AssetCategory

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAssetCategory error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
