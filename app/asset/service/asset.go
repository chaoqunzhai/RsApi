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

type Asset struct {
	service.Service
}

// GetPage 获取Asset列表
func (e *Asset) GetPage(c *dto.AssetGetPageReq, p *actions.DataPermission, list *[]models.Asset, count *int64) error {
	var err error
	var data models.Asset

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AssetService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Asset对象
func (e *Asset) Get(d *dto.AssetGetReq, p *actions.DataPermission, model *models.Asset) error {
	var data models.Asset

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAsset error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Asset对象
func (e *Asset) Insert(c *dto.AssetInsertReq) error {
	var err error
	var data models.Asset
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AssetService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Asset对象
func (e *Asset) Update(c *dto.AssetUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Asset{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AssetService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Asset
func (e *Asset) Remove(d *dto.AssetDeleteReq, p *actions.DataPermission) error {
	var data models.Asset

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAsset error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
