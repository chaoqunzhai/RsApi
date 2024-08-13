package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type RsContract struct {
	service.Service
}

// GetPage 获取RsContract列表
func (e *RsContract) GetPage(c *dto.RsContractGetPageReq, p *actions.DataPermission, list *[]models.RsContract, count *int64) error {
	var err error
	var data models.RsContract

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsContractService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsContract对象
func (e *RsContract) Get(d *dto.RsContractGetReq, p *actions.DataPermission, model *models.RsContract) error {
	var data models.RsContract

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsContract error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RsContract对象
func (e *RsContract) Insert(c *dto.RsContractInsertReq) error {
    var err error
    var data models.RsContract
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsContractService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RsContract对象
func (e *RsContract) Update(c *dto.RsContractUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.RsContract{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("RsContractService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除RsContract
func (e *RsContract) Remove(d *dto.RsContractDeleteReq, p *actions.DataPermission) error {
	var data models.RsContract

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveRsContract error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
