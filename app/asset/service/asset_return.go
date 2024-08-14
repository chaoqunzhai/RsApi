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

type AssetReturn struct {
	service.Service
}

// GetPage 获取AssetReturn列表
func (e *AssetReturn) GetPage(c *dto.AssetReturnGetPageReq, p *actions.DataPermission, list *[]models.AssetReturn, count *int64) error {
	var err error
	var data models.AssetReturn

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AssetReturnService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AssetReturn对象
func (e *AssetReturn) Get(d *dto.AssetReturnGetReq, p *actions.DataPermission, model *models.AssetReturn) error {
	var data models.AssetReturn

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAssetReturn error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AssetReturn对象
func (e *AssetReturn) Insert(c *dto.AssetReturnInsertReq) error {
	var err error
	var data models.AssetReturn
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AssetReturnService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AssetReturn对象
func (e *AssetReturn) Update(c *dto.AssetReturnUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AssetReturn{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AssetReturnService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AssetReturn
func (e *AssetReturn) Remove(d *dto.AssetReturnDeleteReq, p *actions.DataPermission) error {
	var data models.AssetReturn

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAssetReturn error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
