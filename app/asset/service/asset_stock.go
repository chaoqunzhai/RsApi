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

type AssetStock struct {
	service.Service
}

// GetPage 获取AssetStock列表
func (e *AssetStock) GetPage(c *dto.AssetStockGetPageReq, p *actions.DataPermission, list *[]models.AssetStock, count *int64) error {
	var err error
	var data models.AssetStock

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AssetStockService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AssetStock对象
func (e *AssetStock) Get(d *dto.AssetStockGetReq, p *actions.DataPermission, model *models.AssetStock) error {
	var data models.AssetStock

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAssetStock error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AssetStock对象
func (e *AssetStock) Insert(c *dto.AssetStockInsertReq) error {
	var err error
	var data models.AssetStock
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AssetStockService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AssetStock对象
func (e *AssetStock) Update(c *dto.AssetStockUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AssetStock{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AssetStockService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AssetStock
func (e *AssetStock) Remove(d *dto.AssetStockDeleteReq, p *actions.DataPermission) error {
	var data models.AssetStock

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAssetStock error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
