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

type AssetSupplier struct {
	service.Service
}

// GetPage 获取AssetSupplier列表
func (e *AssetSupplier) GetPage(c *dto.AssetSupplierGetPageReq, p *actions.DataPermission, list *[]models.AssetSupplier, count *int64) error {
	var err error
	var data models.AssetSupplier

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AssetSupplierService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AssetSupplier对象
func (e *AssetSupplier) Get(d *dto.AssetSupplierGetReq, p *actions.DataPermission, model *models.AssetSupplier) error {
	var data models.AssetSupplier

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAssetSupplier error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AssetSupplier对象
func (e *AssetSupplier) Insert(c *dto.AssetSupplierInsertReq) error {
	var err error
	var data models.AssetSupplier
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AssetSupplierService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AssetSupplier对象
func (e *AssetSupplier) Update(c *dto.AssetSupplierUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AssetSupplier{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AssetSupplierService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AssetSupplier
func (e *AssetSupplier) Remove(d *dto.AssetSupplierDeleteReq, p *actions.DataPermission) error {
	var data models.AssetSupplier

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAssetSupplier error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
