package service

import (
	"errors"
	"fmt"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type RsDial struct {
	service.Service
}

// GetPage 获取RsDial列表
func (e *RsDial) GetPage(c *dto.RsDialGetPageReq, p *actions.DataPermission, list *[]models.RsDial, count *int64) error {
	var err error
	var data models.RsDial

	var hostIds []int
	orm := e.Orm.Model(&data)
	if c.Search != "" {
		//过滤拨号
		var hostList []models.RsHost

		likeQ := fmt.Sprintf(" sn like '%%%s%%' or host_name like '%%%s%%' ", c.Search, c.Search)
		e.Orm.Model(&models.RsHost{}).Where(likeQ).Find(&hostList)
		for _, host := range hostList {
			hostIds = append(hostIds, host.Id)
		}
		if len(hostIds) > 0 {
			orm = orm.Where("host_id in ?", hostIds)
		}
	}
	err = orm.Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		actions.Permission(data.TableName(), p),
	).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsDialService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsDial对象
func (e *RsDial) Get(d *dto.RsDialGetReq, p *actions.DataPermission, model *models.RsDial) error {
	var data models.RsDial

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsDial error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	var IdcRow models.RsIdc
	if data.IdcId > 0 {
		e.Orm.Model(&models.RsIdc{}).Where("id = ?", data.IdcId).Find(&IdcRow)
		data.IdcInfo = IdcRow
	}
	return nil
}

// Insert 创建RsDial对象
func (e *RsDial) Insert(c *dto.RsDialInsertReq) error {
	var err error
	var data models.RsDial
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsDialService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RsDial对象
func (e *RsDial) Update(c *dto.RsDialUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RsDial{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RsDialService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RsDial
func (e *RsDial) Remove(d *dto.RsDialDeleteReq, p *actions.DataPermission) error {
	var data models.RsDial

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRsDial error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
