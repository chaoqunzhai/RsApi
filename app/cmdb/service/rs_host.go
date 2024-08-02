package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service/dto"
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
	"go-admin/common/utils"
	"gorm.io/gorm"
	"strings"
)

type RsHost struct {
	service.Service
}

// GetPage 获取RsHost列表
func (e *RsHost) GetPage(c *dto.RsHostGetPageReq, p *actions.DataPermission, list *[]models.RsHost, count *int64) error {
	var err error
	var data models.RsHost

	orm := e.Orm.Model(&data)
	if c.IdcName != "" {
		var idcList []models.RsIdc
		e.Orm.Model(&models.RsIdc{}).Select("id").Where("name like ?", fmt.Sprintf("%%%v%%", c.IdcName)).Find(&idcList)
		var cache []int
		for _, idc := range idcList {
			cache = append(cache, idc.Id)
		}
		orm = orm.Where("idc in (?)", cache)
	}
	if c.IdcNumber != "" {
		var idcList []models.RsIdc
		e.Orm.Model(&models.RsIdc{}).Select("id").Where("number like ?", fmt.Sprintf("%%%v%%", c.IdcName)).Find(&idcList)
		var cache []int
		for _, idc := range idcList {
			cache = append(cache, idc.Id)
		}
		orm = orm.Where("idc in (?)", cache)
	}

	if c.BusinessId != "" {
		var bindHostId []int

		e.Orm.Raw(fmt.Sprintf("select host_id from host_bind_business where business_id in (%v)", c.BusinessId)).Scan(&bindHostId)

		orm = orm.Where("id in (?)", bindHostId)
	}
	err = orm.Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		actions.Permission(data.TableName(), p),
	).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsHostService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsHost对象
func (e *RsHost) Get(d *dto.RsHostGetReq, p *actions.DataPermission, model *models.RsHost) error {
	var data models.RsHost

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Preload("Business").Preload("Idc").Preload("Tag").
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsHost error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RsHost对象
func (e *RsHost) Insert(c *dto.RsHostInsertReq) error {
	var err error
	var data models.RsHost
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsHostService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RsHost对象
func (e *RsHost) Update(c *dto.RsHostUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RsHost{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RsHostService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RsHost
func (e *RsHost) Remove(d *dto.RsHostDeleteReq, p *actions.DataPermission) error {
	var data models.RsHost

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRsHost error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
func (e *RsHost) GetIdcList(ids []int) map[int][]interface{} {

	RowMap := make(map[int][]interface{}, 0)

	ids = utils.RemoveRepeatInt(ids)
	var RowList []models.RsIdc
	e.Orm.Model(&models.RsIdc{}).Where("id in ? ", ids).Find(&RowList)

	for _, idc := range RowList {
		cache, ok := RowMap[idc.Id]
		if !ok {
			cache = make([]interface{}, 0)
		}
		dat := map[string]interface{}{
			"id":      idc.Id,
			"name":    idc.Name,
			"number":  idc.Number,
			"region":  idc.Region,
			"address": idc.Address,
		}
		cache = append(cache, dat)
		RowMap[idc.Id] = cache
	}
	return RowMap

}
func (e *RsHost) GetHostSoftware(ids []int) map[int][]models2.HostSoftware {
	HostSoftwareMap := make(map[int][]models2.HostSoftware, 0)

	var HostSoftwareList []models2.HostSoftware
	e.Orm.Model(&models2.HostSoftware{}).Where("host_id in ? ", ids).Find(&HostSoftwareList)

	for _, row := range HostSoftwareList {
		cache, ok := HostSoftwareMap[row.HostId]
		if !ok {
			cache = make([]models2.HostSoftware, 0)
		}
		cache = append(cache, row)
		HostSoftwareMap[row.HostId] = cache
	}
	return HostSoftwareMap

}
func (e *RsHost) GetMonitorData(ids []int) map[int][]interface{} {
	var monitorList []models2.HostSystem

	result := make(map[int][]interface{}, 0)
	e.Orm.Model(&models2.HostSystem{}).Where("host_id in ?", ids).Find(&monitorList)

	if len(monitorList) == 0 {
		return map[int][]interface{}{}
	}

	for _, row := range monitorList {
		dat := map[string]interface{}{
			"transmit": row.TransmitNumber,
			"receive":  row.ReceiveNumber,
		}

		if row.MemoryData != "" {

			HostMemory := dto.HostMemory{}
			if unErr := json.Unmarshal([]byte(row.MemoryData), &HostMemory); unErr != nil {
				continue

			}
			dat["used_percent"] = fmt.Sprintf("%.2f", 100*float64(HostMemory.U)/float64(HostMemory.T))      //已使用百分比
			dat["available_percent"] = fmt.Sprintf("%.2f", 100*float64(HostMemory.A)/float64(HostMemory.T)) //可使用百分比
		}

		cache, ok := result[row.HostId]
		if !ok {
			cache = make([]interface{}, 0)
		}
		cache = append(cache, dat)
		result[row.Id] = cache
	}

	return result
}

func (e *RsHost) GetCity(row models.RsHost) string {

	return ""
}

func (e *RsHost) GetBusinessMap(ids []int) map[int][]dto.LabelRow {

	result := make(map[int][]dto.LabelRow, 0)
	var RsBusinessList []models.RsBusiness
	var cacheIds []string
	for _, i := range ids {
		cacheIds = append(cacheIds, fmt.Sprintf("%v", i))
	}
	hostBindIdc := fmt.Sprintf("select business_id,host_id from host_bind_business where `host_id` in (%v)", strings.Join(cacheIds, ","))
	var bindIds []struct {
		HostId     int `json:"host_id"`
		BusinessId int `json:"business_id"`
	}
	e.Orm.Raw(hostBindIdc).Scan(&bindIds)

	if len(bindIds) == 0 {

		return map[int][]dto.LabelRow{}
	}

	var cacheBuIds []int
	for _, r := range bindIds {
		cacheBuIds = append(cacheBuIds, r.BusinessId)
	}
	e.Orm.Model(&models.RsBusiness{}).Select("name,id").Where("id in ?", cacheBuIds).Find(&RsBusinessList)

	BusinessMap := make(map[int]dto.LabelRow, 0)
	for _, b := range RsBusinessList {
		LabelRow := dto.LabelRow{
			Id:    b.Id,
			Label: b.Name,
			Value: fmt.Sprintf("%v", b.Id),
		}
		BusinessMap[b.Id] = LabelRow
	}
	for _, row := range bindIds {
		buDat, buOk := BusinessMap[row.BusinessId]
		if !buOk {
			continue
		}
		cache, ok := result[row.HostId]
		if !ok {
			cache = make([]dto.LabelRow, 0)
		}
		cache = append(cache, buDat)
		result[row.HostId] = cache

	}

	return result
}
