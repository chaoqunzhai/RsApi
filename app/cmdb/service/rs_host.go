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
		e.Orm.Model(&models.RsIdc{}).Select("id").Where("number like ?", fmt.Sprintf("%%%v%%", c.IdcNumber)).Find(&idcList)
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

		fmt.Println("查询业务", bindHostId, len(bindHostId))
	}

	if c.HostName != "" {
		//批量把\n换成逗号
		newHostName := strings.Replace(c.HostName, "\n", ",", -1)
		// 批量把空格换成逗号
		newHostName = strings.Replace(newHostName, " ", ",", -1)

		//一个元素 是模糊搜索
		newHostList := strings.Split(newHostName, ",")
		if len(newHostList) == 1 {
			likeKey := fmt.Sprintf("%%%v%%", newHostName)
			orm = orm.Where("host_name like ? OR sn like ?", likeKey, likeKey)
		} else {
			//多个元素 就是精确搜索了
			orm = orm.Where("host_name in ? OR sn in ?", newHostList, newHostList)
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
		).Preload("Business").Preload("Tag").First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsHost error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}

	//做一下扩充字段补齐

	var NetDevice []models2.HostNetDevice
	e.Orm.Model(&models2.HostNetDevice{}).Where("host_id = ?", d.Id).Find(&NetDevice)

	var DialList []models.RsDial
	e.Orm.Model(&models.RsDial{}).Where("host_id = ?", d.Id).Find(&DialList)

	system := map[string]interface{}{
		"cpu": model.Cpu,
		"ip":  model.Ip,
		"memory": func() int {
			if model.Memory == 0 {
				return 0
			}
			return int(model.Memory / 1024 / 1024 / 1024)
		}(),
		"kernel": model.Kernel,
	}

	ids := []int{d.Id}

	HostMapMonitorData := e.GetMonitorData(ids)

	IdcMapData := e.GetIdcList([]int{model.Idc})

	ExtendHostInfo := models.ExtendHostInfo{
		NetDevice: NetDevice,
		System:    system,
		DialList:  DialList,
	}
	if dat, ok := HostMapMonitorData[d.Id]; ok {

		ExtendHostInfo.MemoryMonitor = dat["memory"]
		ExtendHostInfo.Disk = dat["disk"]
	}
	if idcInfo, ok := IdcMapData[model.Idc]; ok {
		if len(idcInfo) > 0 {
			model.IdcInfo = idcInfo[0]
		}
	}

	model.ExtendHostInfo = ExtendHostInfo
	return nil
}

// Insert 创建RsHost对象
func (e *RsHost) Insert(c *dto.RsHostInsertReq) (id int, err error) {

	var data models.RsHost
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsHostService Insert error:%s \r\n", err)
		return 0, err
	}
	return data.Id, nil
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
	e.Orm.Model(models2.HostSystem{}).Where("host_id in ?", d.GetId()).Unscoped().Delete(&models2.HostSystem{})
	e.Orm.Model(models2.HostSwitchLog{}).Where("host_id in ?", d.GetId()).Unscoped().Delete(&models2.HostSwitchLog{})
	e.Orm.Model(models2.HostSoftware{}).Where("host_id in ?", d.GetId()).Unscoped().Delete(&models2.HostSoftware{})
	e.Orm.Model(models2.HostNetDevice{}).Where("host_id in ?", d.GetId()).Unscoped().Delete(&models2.HostNetDevice{})

	idsStr := d.GetIdStr()
	e.Orm.Exec(fmt.Sprintf("update rs_dial  set deleted_at = NULL where host_id in (%v)", idsStr))

	e.Orm.Exec(fmt.Sprintf("DELETE from host_bind_business where host_id in (%v)", idsStr))
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

type DialCnfRow struct {
	AllLine  int    `json:"allLine"`
	DialUp   int    `json:"dialUp"`
	DialDown int    `json:"dialDown"`
	NetUp    int    `json:"netUp"`
	NetDown  int    `json:"netDown"`
	Info     string `json:"info"`
}

func (e *RsHost) GetDialData(ids []int) map[int]DialCnfRow {
	var dialList []models.RsDial
	e.Orm.Model(&models.RsDial{}).Where("host_id in ? ", ids).Find(&dialList)

	result := make(map[int]DialCnfRow, 0)

	for _, row := range dialList {
		data, ok := result[row.HostId]
		if !ok {
			data = DialCnfRow{}
		}
		data.AllLine += 1

		if row.Status == 1 { //拨通了

			data.DialUp += 1
		} else {

			data.DialDown += 1
		}
		if row.NetworkingStatus == 1 {

			data.NetUp += 1
		} else {
			data.NetDown += 1
		}

		result[row.HostId] = data
	}

	return result
}

func (e *RsHost) GetMonitorData(ids []int) map[int]map[string]interface{} {
	var monitorList []models2.HostSystem

	result := make(map[int]map[string]interface{}, 0)
	e.Orm.Model(&models2.HostSystem{}).Where("host_id in ?", ids).Find(&monitorList)

	if len(monitorList) == 0 {
		return map[int]map[string]interface{}{}
	}

	for _, row := range monitorList {
		dat := map[string]interface{}{
			"transmit": row.TransmitNumber,
			"receive":  row.ReceiveNumber,
		}
		cache, ok := result[row.HostId]
		if !ok {
			cache = make(map[string]interface{}, 0)
		}
		if row.MemoryData != "" {

			HostMemory := dto.HostMemory{}
			if unErr := json.Unmarshal([]byte(row.MemoryData), &HostMemory); unErr == nil {
				dat["used_percent"] = fmt.Sprintf("%.2f", 100*float64(HostMemory.U)/float64(HostMemory.T))      //已使用百分比
				dat["available_percent"] = fmt.Sprintf("%.2f", 100*float64(HostMemory.A)/float64(HostMemory.T)) //可使用百分比
			}
		}

		cache["memory"] = dat
		HostDisk := make([]dto.HDDevUsage, 0)
		if row.Disk != "" {

			_ = json.Unmarshal([]byte(row.Disk), &HostDisk)

		}
		cache["disk"] = HostDisk
		result[row.HostId] = cache
	}

	return result
}

func (e *RsHost) GetCity(row models.RsHost) string {

	return ""
}

func GetHostBindBusinessMap(orm *gorm.DB, ids []int) map[int][]dto.LabelRow {

	result := make(map[int][]dto.LabelRow, 0)
	var RsBusinessList []models.RsBusiness
	var cacheIds []string
	for _, i := range ids {
		cacheIds = append(cacheIds, fmt.Sprintf("%v", i))
	}
	if len(cacheIds) == 0 {
		return map[int][]dto.LabelRow{}
	}
	hostBindIdc := fmt.Sprintf("select business_id,host_id from host_bind_business where `host_id` in (%v)", strings.Join(cacheIds, ","))
	var bindIds []struct {
		HostId     int `json:"host_id"`
		BusinessId int `json:"business_id"`
	}
	orm.Raw(hostBindIdc).Scan(&bindIds)

	if len(bindIds) == 0 {

		return map[int][]dto.LabelRow{}
	}
	var cacheBuIds []int
	for _, r := range bindIds {
		cacheBuIds = append(cacheBuIds, r.BusinessId)
	}
	orm.Model(&models.RsBusiness{}).Select("name,id").Where("id in ?", cacheBuIds).Find(&RsBusinessList)

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
