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

type AssetRecording struct {
	service.Service
}

// GetPage 获取AssetRecording列表
func (e *AssetRecording) GetPage(c *dto.AssetRecordingGetPageReq, p *actions.DataPermission, list *[]models.AssetRecording, count *int64) error {
	var err error
	var data models.AssetRecording

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AssetRecordingService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AssetRecording对象
func (e *AssetRecording) Get(d *dto.AssetRecordingGetReq, p *actions.DataPermission, model *models.AssetRecording) error {
	var data models.AssetRecording

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAssetRecording error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AssetRecording对象
func (e *AssetRecording) Insert(c *dto.AssetRecordingInsertReq) error {
    var err error
    var data models.AssetRecording
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AssetRecordingService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改AssetRecording对象
func (e *AssetRecording) Update(c *dto.AssetRecordingUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.AssetRecording{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("AssetRecordingService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除AssetRecording
func (e *AssetRecording) Remove(d *dto.AssetRecordingDeleteReq, p *actions.DataPermission) error {
	var data models.AssetRecording

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveAssetRecording error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
