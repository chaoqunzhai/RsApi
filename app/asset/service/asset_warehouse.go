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

type AssetWarehouse struct {
	service.Service
}

// GetPage 获取AssetWarehouse列表
func (e *AssetWarehouse) GetPage(c *dto.AssetWarehouseGetPageReq, p *actions.DataPermission, list *[]models.AssetWarehouse, count *int64) error {
	var err error
	var data models.AssetWarehouse

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AssetWarehouseService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AssetWarehouse对象
func (e *AssetWarehouse) Get(d *dto.AssetWarehouseGetReq, p *actions.DataPermission, model *models.AssetWarehouse) error {
	var data models.AssetWarehouse

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAssetWarehouse error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AssetWarehouse对象
func (e *AssetWarehouse) Insert(c *dto.AssetWarehouseInsertReq) error {
	var err error
	var data models.AssetWarehouse
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AssetWarehouseService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AssetWarehouse对象
func (e *AssetWarehouse) Update(c *dto.AssetWarehouseUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AssetWarehouse{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AssetWarehouseService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AssetWarehouse
func (e *AssetWarehouse) Remove(d *dto.AssetWarehouseDeleteReq, p *actions.DataPermission) error {
	var data models.AssetWarehouse

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAssetWarehouse error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
