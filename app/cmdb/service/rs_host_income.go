package service

import (
	"errors"
	"strings"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type RsHostIncome struct {
	service.Service
}

// GetPage 获取RsHostIncome列表
func (e *RsHostIncome) GetPage(c *dto.RsHostIncomeGetPageReq, p *actions.DataPermission, list *[]models.RsHostIncome, count *int64) error {
	var err error
	var data models.RsHostIncome

	orm := e.Orm.Model(&data)

	if c.BusinessId != "" {
		orm = orm.Where("bu_id in ?", strings.Split(c.BusinessId, ","))
	}
	if c.HostSearch != "" {

		e.Orm.Model(&models.RsHost{}).Where("host_name like ? or sn like ?", "%"+c.HostSearch+"%", "%"+c.HostSearch+"%")
	}
	if c.Region != "" {

	}
	if c.IncomeMonth != "" {
		orm = orm.Where("alg_day like ? ",c.IncomeMonth+"%" )
	}
	if c.CustomId != "" {
		//先通过客户搜索机房
		//搜索到的机房 在收益中查询
		idcList := make([]string, 0)
		e.Orm.Model(&models.RsIdc{}).Select("id").Where("custom_id = ?", c.CustomId).Find(&idcList).Scan(&idcList)
		orm = orm.Where("idc_id in ?", idcList)
	}
	err = orm.Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		actions.Permission(data.TableName(), p),
	).Order("alg_day desc").
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RsHostIncomeService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RsHostIncome对象
func (e *RsHostIncome) Get(d *dto.RsHostIncomeGetReq, p *actions.DataPermission, model *models.RsHostIncome) error {
	var data models.RsHostIncome

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsHostIncome error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RsHostIncome对象
func (e *RsHostIncome) Insert(c *dto.RsHostIncomeInsertReq) error {
	var err error
	var data models.RsHostIncome
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsHostIncomeService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RsHostIncome对象
func (e *RsHostIncome) Update(c *dto.RsHostIncomeUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RsHostIncome{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RsHostIncomeService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RsHostIncome
func (e *RsHostIncome) Remove(d *dto.RsHostIncomeDeleteReq, p *actions.DataPermission) error {
	var data models.RsHostIncome

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRsHostIncome error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
