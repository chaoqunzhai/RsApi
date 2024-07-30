package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	models2 "go-admin/cmd/migrate/migration/models"
	"gorm.io/gorm"

	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type RsHost struct {
	service.Service
}

// GetPage 获取RsHost列表
func (e *RsHost) GetPage(c *dto.RsHostGetPageReq, p *actions.DataPermission, list *[]models.RsHost, count *int64) error {
	var err error
	var data models.RsHost

	err = e.Orm.Model(&data).
		Scopes(
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

func (e *RsHost) GetMonitorData(row models.RsHost) map[string]interface{} {
	var monitor models2.HostSystem

	result := make(map[string]interface{}, 0)
	e.Orm.Model(&models2.HostSystem{}).Where("host_id = ?", row.Id).Limit(1).Find(&monitor)

	if monitor.Id == 0 {
		return result
	}

	result = map[string]interface{}{
		"transmit": monitor.TransmitNumber,
		"receive":  monitor.ReceiveNumber,
	}

	if monitor.MemoryData != "" {

		HostMemory := dto.HostMemory{}
		if unErr := json.Unmarshal([]byte(monitor.MemoryData), &HostMemory); unErr != nil {
			return result

		}
		result["used_percent"] = fmt.Sprintf("%.2f", 100*float64(HostMemory.U)/float64(HostMemory.T))      //已使用百分比
		result["available_percent"] = fmt.Sprintf("%.2f", 100*float64(HostMemory.A)/float64(HostMemory.T)) //可使用百分比
	}

	return result
}

func (e *RsHost) GetIdcInfo(row models.RsHost) map[string]interface{} {

	var idc models.RsIdc

	if row.Idc == 0 {

		return map[string]interface{}{
			"id":     0,
			"name":   "",
			"number": "",
			"region": "",
		}
	}
	e.Orm.Model(&idc).Select("name,id").Where("id = ?", row.Idc).Limit(1).Find(&idc)

	return map[string]interface{}{
		"id":     idc.Id,
		"name":   idc.Name,
		"number": idc.Number,
		"region": idc.Region,
	}
}
func (e *RsHost) GetCity(row models.RsHost) string {

	return ""
}

func (e *RsHost) GetBusiness(row models.RsHost) []interface{} {

	list := make([]interface{}, 0)
	var RsBusinessList []models.RsBusiness
	hostBindIdc := fmt.Sprintf("select business_id from host_bind_business where `host_id` = %v", row.Id)
	var bindIds []interface{}
	e.Orm.Raw(hostBindIdc).Scan(&bindIds)

	if len(bindIds) == 0 {

		return list
	}
	e.Orm.Model(&models.RsBusiness{}).Select("name,id").Where("id in ?", bindIds).Find(&RsBusinessList)

	for _, b := range RsBusinessList {

		list = append(list, dto.LabelRow{
			Id:    b.Id,
			Label: b.Name,
			Value: fmt.Sprintf("%v", b.Id),
		})
	}
	return list
}
