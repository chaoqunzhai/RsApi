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

type Combination struct {
	service.Service
}

// GetPage 获取Combination列表
func (e *Combination) GetPage(c *dto.CombinationGetPageReq, p *actions.DataPermission, list *[]models.Combination, count *int64) error {
	var err error
	var data models.Combination

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("CombinationService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Combination对象
func (e *Combination) Get(d *dto.CombinationGetReq, p *actions.DataPermission, model *models.Combination) error {
	var data models.Combination

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetCombination error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Combination对象
func (e *Combination) Insert(c *dto.CombinationInsertReq) (uid int, err error) {
	var data models.Combination
	c.Generate(&data)
	data.Status = 1
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("CombinationService Insert error:%s \r\n", err)
		return 0, err
	}
	return data.Id, nil
}

// Update 修改Combination对象
func (e *Combination) Update(c *dto.CombinationUpdateReq, p *actions.DataPermission) (uid int, err error) {

	var data = models.Combination{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)
	data.Status = 1
	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("CombinationService Save error:%s \r\n", err)
		return 0, err
	}
	if db.RowsAffected == 0 {
		return 0, errors.New("无权更新该数据")
	}
	return data.Id, nil
}

// Remove 删除Combination
func (e *Combination) Remove(newIds []int, p *actions.DataPermission) error {
	var data models.Combination

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, newIds)
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveCombination error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
