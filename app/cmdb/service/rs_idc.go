package service

import (
	"errors"
	"fmt"
	"go-admin/common/utils"
	"go-admin/global"
	"strings"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type RsIdc struct {
	service.Service
}

// GetPage 获取RsIdc列表
func (e *RsIdc) GetPage(c *dto.RsIdcGetPageReq, p *actions.DataPermission, list *[]models.RsIdc, count *int64) error {
	var err error
	var data models.RsIdc

	orm := e.Orm.Model(&data)
	if c.Search != "" {

		likeQ := fmt.Sprintf("number like '%%%s%%' or name like '%%%s%%' ", c.Search, c.Search)
		orm = orm.Where(likeQ)
	}
	if c.CustomId != "" {
		if c.CustomId == "empty" {
			orm = orm.Where("custom_id = 0 OR custom_id IS  NULL")
		} else {
			orm = orm.Where("custom_id = ?", c.CustomId)
		}
	}
	if c.Region != "" {
		RegionList := strings.Split(c.Region, ",")
		var searchRegion string

		if len(RegionList) > 1 {

			searchRegion = RegionList[len(RegionList)-1]
		} else {
			searchRegion = c.Region
		}
		likeQ := fmt.Sprintf("region like '%%%s%%'", searchRegion)
		orm = orm.Where(likeQ)
	}

	if c.OffLineOrder != "asc" {
		//离线数量排序

		var hostList []models.RsHost
		//1.查询所有的主机
		e.Orm.Model(&models.RsHost{}).Select("idc,status").Where("status = ?", global.HostOffline).Find(&hostList)
		idcMap := make(map[int]int, 0)

		for _, host := range hostList {
			RsIdcCount, ok := idcMap[host.Idc]
			if !ok {
				RsIdcCount = 0
			}

			RsIdcCount += 1
			idcMap[host.Idc] = RsIdcCount
		}

		kvSlice := utils.SortMap(idcMap, c.OffLineOrder)

		idcS := func() []int {
			var idcList []int
			for _, v := range kvSlice {
				idcList = append(idcList, v.Key)
			}
			return idcList
		}()

		if len(idcS) > 0 {
			orm = orm.Where("id in (?)", idcS)
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
		e.Log.Errorf("RsIdcService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsIdc对象
func (e *RsIdc) Get(d *dto.RsIdcGetReq, p *actions.DataPermission, model *models.RsIdc) error {
	var data models.RsIdc

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsIdc error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RsIdc对象
func (e *RsIdc) Insert(c *dto.RsIdcInsertReq) (id int, err error) {

	var data models.RsIdc
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsIdcService Insert error:%s \r\n", err)
		return 0, err
	}
	return data.Id, nil
}

// Update 修改RsIdc对象
func (e *RsIdc) Update(c *dto.RsIdcUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RsIdc{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RsIdcService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RsIdc
func (e *RsIdc) Remove(d *dto.RsIdcDeleteReq, p *actions.DataPermission) error {
	var data models.RsIdc

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRsIdc error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
