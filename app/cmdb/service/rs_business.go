package service

import (
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

type RsBusiness struct {
	service.Service
}

func GetBusinessMap(orm *gorm.DB, ids []int) map[int]*models.RsBusiness {

	var RsBusinessList []*models.RsBusiness

	orm.Model(&models.RsBusiness{}).Where("id in ?", ids).Find(&RsBusinessList)

	BusinessMap := make(map[int]*models.RsBusiness, 0)
	for _, b := range RsBusinessList {

		BusinessMap[b.Id] = b
	}

	return BusinessMap
}

func (e *RsBusiness) GetChildren(parentId int) interface{} {

	var list []models.RsBusiness
	e.Orm.Model(&models.RsBusiness{}).Where("parent_id = ?", parentId).Find(&list)

	return list

}

// GetPage 获取RsBusiness列表
func (e *RsBusiness) GetPage(c *dto.RsBusinessGetPageReq, p *actions.DataPermission, list *[]models.RsBusiness, count *int64) error {
	var err error
	var data models.RsBusiness
	orm := e.Orm.Model(&data)
	if c.TreeTag > 0 {
		orm = orm.Where("parent_id = 0 OR parent_id is NULL")
	}
	err = orm.Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		actions.Permission(data.TableName(), p),
	).Find(list).Limit(-1).Offset(-1).
		Count(count).Error

	if err != nil {
		e.Log.Errorf("RsBusinessService GetPage error:%s \r\n", err)
		return err
	}

	return nil
}

// Get 获取RsBusiness对象
func (e *RsBusiness) Get(d *dto.RsBusinessGetReq, p *actions.DataPermission, model *models.RsBusiness) error {
	var data models.RsBusiness

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRsBusiness error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	var costList []*models.RsBusinessCostCnf

	e.Orm.Model(&models.RsBusinessCostCnf{}).Where("bu_id = ?", d.GetId()).Find(&costList)

	var Children []*models.RsBusiness
	e.Orm.Model(&models.RsBusiness{}).Where("parent_id = ?", d.GetId()).Find(&Children)
	model.Children = Children
	model.CostCnf = costList

	return nil
}

// Insert 创建RsBusiness对象
func (e *RsBusiness) Insert(c *dto.RsBusinessInsertReq) (id int, err error) {

	var data models.RsBusiness
	c.Generate(&data)
	data.Layer = 1
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RsBusinessService Insert error:%s \r\n", err)
		return 0, err
	}
	for _, row := range c.CostCnf {
		var costCnf models.RsBusinessCostCnf
		row.Generate(&costCnf)
		costCnf.BuId = data.Id
		_ = e.Orm.Create(&costCnf)
	}
	return data.Id, nil
}

// Update 修改RsBusiness对象
func (e *RsBusiness) Update(CreateUser string, c *dto.RsBusinessUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RsBusiness{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RsBusinessService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}

	hasIds := make([]int, 0)
	var bindCostCnf []models.RsBusinessCostCnf
	e.Orm.Model(&models.RsBusinessCostCnf{}).Select("id").Where("bu_id = ?", c.Id).Find(&bindCostCnf)

	for _, row := range bindCostCnf {
		hasIds = append(hasIds, row.Id)
	}

	updateId := make([]int, 0)
	var diffIds []int
	if len(c.CostCnf) == 0 {
		diffIds = hasIds
	} else {
		for _, row := range c.CostCnf {
			var infoList []string

			var bandRow models.RsBusinessCostCnf
			if row.Id > 0 { //更新
				e.Orm.Model(&bandRow).First(&bandRow, row.GetId())

				fmt.Println("价格比较", bandRow.Price, row.Price)
				if bandRow.Price != row.Price {
					infoList = append(infoList, fmt.Sprintf("计算价格变更 %v->%v", row.Price, bandRow.Price))
				}
				row.Generate(&bandRow)
				bandRow.BuId = data.Id
				updateId = append(updateId, row.Id)

				//if bandRow.Start.Format(time.DateTime) != row.Start {
				//	infoList = append(infoList, fmt.Sprintf("开始区间变更 %v->%v", row.Start))
				//}
				//if bandRow.End.Format(time.DateTime) != row.End {
				//	infoList = append(infoList, fmt.Sprintf("结束区间变更  %v->%v", row.End))
				//}
				//if bandRow.RangePrice != row.RangePrice {
				//	infoList = append(infoList, fmt.Sprintf("区间价格变更  %v->%v", row.Price, row.RangePrice))
				//}

				e.Orm.Save(&bandRow)
			} else { //创建
				row.Generate(&bandRow)
				bandRow.BuId = data.Id
				err = e.Orm.Create(&bandRow).Error
				//if row.Start != "" {
				//	infoList = append(infoList, fmt.Sprintf("设置开始区间 %v", row.Start))
				//}
				//if row.End != "" {
				//	infoList = append(infoList, fmt.Sprintf("设置结束区间 %v", row.End))
				//}
				if row.RangePrice > 0 {
					infoList = append(infoList, fmt.Sprintf("设置区间价格 %v", row.RangePrice))
				}
			}

			e.Orm.Create(&models2.OperationLog{
				CreateUser: CreateUser,
				Module:     "rs_business",
				Action:     "PUT",
				ObjectId:   data.Id,
				TargetId:   data.Id,
				Info:       strings.Join(infoList, " "),
			})
		}
		diffIds = utils.DifferenceInt(hasIds, updateId)
	}

	if len(diffIds) > 0 {
		e.Orm.Model(&models.RsBusinessCostCnf{}).Where("id in ?", diffIds).Delete(&models.RsBusinessCostCnf{})
	}

	return nil
}

// Remove 删除RsBusiness
func (e *RsBusiness) Remove(d *dto.RsBusinessDeleteReq, p *actions.DataPermission) error {
	var data models.RsBusiness

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRsBusiness error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}

	//删除对应的带宽配置信息
	var bindCost models.RsBusinessCostCnf
	e.Orm.Model(&bindCost).Delete(&bindCost, d.GetId())
	return nil
}
