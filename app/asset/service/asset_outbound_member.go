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

type AssetOutboundMember struct {
	service.Service
}

// GetPage 获取AssetOutboundMember列表
func (e *AssetOutboundMember) GetPage(c *dto.AssetOutboundMemberGetPageReq, p *actions.DataPermission, list *[]models.AssetOutboundMember, count *int64) error {
	var err error
	var data models.AssetOutboundMember

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AssetOutboundMemberService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AssetOutboundMember对象
func (e *AssetOutboundMember) Get(d *dto.AssetOutboundMemberGetReq, p *actions.DataPermission, model *models.AssetOutboundMember) error {
	var data models.AssetOutboundMember

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAssetOutboundMember error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AssetOutboundMember对象
func (e *AssetOutboundMember) Insert(c *dto.AssetOutboundMemberInsertReq) error {
	var err error
	var data models.AssetOutboundMember
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AssetOutboundMemberService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AssetOutboundMember对象
func (e *AssetOutboundMember) Update(c *dto.AssetOutboundMemberUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AssetOutboundMember{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AssetOutboundMemberService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AssetOutboundMember
func (e *AssetOutboundMember) Remove(d *dto.AssetOutboundMemberDeleteReq, p *actions.DataPermission) error {
	var data models.AssetOutboundMember

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAssetOutboundMember error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
