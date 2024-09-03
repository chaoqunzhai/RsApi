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

type AssetDisposal struct {
	service.Service
}

// GetPage 获取AssetDisposal列表
func (e *AssetDisposal) GetPage(c *dto.AssetDisposalGetPageReq, p *actions.DataPermission, list *[]models.AssetDisposal, count *int64) error {
	var err error
	var data models.AssetDisposal

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AssetDisposalService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AssetDisposal对象
func (e *AssetDisposal) Get(d *dto.AssetDisposalGetReq, p *actions.DataPermission, model *models.AssetDisposal) error {
	var data models.AssetDisposal

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAssetDisposal error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AssetDisposal对象
func (e *AssetDisposal) Insert(c *dto.AssetDisposalInsertReq) error {
	var err error
	var data models.AssetDisposal
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AssetDisposalService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AssetDisposal对象
func (e *AssetDisposal) Update(c *dto.AssetDisposalUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AssetDisposal{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AssetDisposalService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AssetDisposal
func (e *AssetDisposal) Remove(d *dto.AssetDisposalDeleteReq, p *actions.DataPermission) error {
	var data models.AssetDisposal

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAssetDisposal error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
