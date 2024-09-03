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

type AssetInboundMember struct {
	service.Service
}

// GetPage 获取AssetInboundMember列表
func (e *AssetInboundMember) GetPage(c *dto.AssetInboundMemberGetPageReq, p *actions.DataPermission, list *[]models.AssetInboundMember, count *int64) error {
	var err error
	var data models.AssetInboundMember

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AssetInboundMemberService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AssetInboundMember对象
func (e *AssetInboundMember) Get(d *dto.AssetInboundMemberGetReq, p *actions.DataPermission, model *models.AssetInboundMember) error {
	var data models.AssetInboundMember

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAssetInboundMember error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AssetInboundMember对象
func (e *AssetInboundMember) Insert(c *dto.AssetInboundMemberInsertReq) error {
	var err error
	var data models.AssetInboundMember
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AssetInboundMemberService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AssetInboundMember对象
func (e *AssetInboundMember) Update(c *dto.AssetInboundMemberUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AssetInboundMember{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AssetInboundMemberService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AssetInboundMember
func (e *AssetInboundMember) Remove(d *dto.AssetInboundMemberDeleteReq, p *actions.DataPermission) error {
	var data models.AssetInboundMember

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAssetInboundMember error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
