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

type AssetGroupMember struct {
	service.Service
}

// GetPage 获取AssetGroupMember列表
func (e *AssetGroupMember) GetPage(c *dto.AssetGroupMemberGetPageReq, p *actions.DataPermission, list *[]models.AssetGroupMember, count *int64) error {
	var err error
	var data models.AssetGroupMember

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AssetGroupMemberService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AssetGroupMember对象
func (e *AssetGroupMember) Get(d *dto.AssetGroupMemberGetReq, p *actions.DataPermission, model *models.AssetGroupMember) error {
	var data models.AssetGroupMember

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAssetGroupMember error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AssetGroupMember对象
func (e *AssetGroupMember) Insert(c *dto.AssetGroupMemberInsertReq) error {
	var err error
	var data models.AssetGroupMember
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AssetGroupMemberService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AssetGroupMember对象
func (e *AssetGroupMember) Update(c *dto.AssetGroupMemberUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AssetGroupMember{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AssetGroupMemberService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AssetGroupMember
func (e *AssetGroupMember) Remove(d *dto.AssetGroupMemberDeleteReq, p *actions.DataPermission) error {
	var data models.AssetGroupMember

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAssetGroupMember error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
