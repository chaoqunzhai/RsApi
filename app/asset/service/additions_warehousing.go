package service

import (
	"database/sql"
	"errors"
	"fmt"
	"go-admin/global"
	"strconv"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/asset/models"
	"go-admin/app/asset/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type AdditionsWarehousing struct {
	service.Service
}

// GetPage 获取AdditionsWarehousing列表
func (e *AdditionsWarehousing) GetPage(combinationId string, c *dto.AdditionsWarehousingGetPageReq, p *actions.DataPermission, list *[]models.AdditionsWarehousing, count *int64) error {
	var err error
	var data models.AdditionsWarehousing

	orm := e.Orm.Model(&data)

	if c.Search != "" {

		orm = orm.Where("sn like ? or code like ?", "%"+c.Search+"%", "%"+c.Search+"%")
	}

	if combinationId != "" {

		if combinationId == "0" {

			orm = orm.Where("combination_id = 0 ")
		} else {
			orm = orm.Where("combination_id != 0 ")
		}
	}
	if c.CategoryId > 0 {
		orm = orm.Where("category_id = ?", c.CategoryId)
	} else {
		orm = orm.Where("category_id != 1")
	}
	err = orm.Scopes(
		cDto.MakeCondition(c.GetNeedSearch()),
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		actions.Permission(data.TableName(), p),
	).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AdditionsWarehousingService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取AdditionsWarehousing对象
func (e *AdditionsWarehousing) Get(d *dto.AdditionsWarehousingGetReq, p *actions.DataPermission, model *models.AdditionsWarehousing) error {
	var data models.AdditionsWarehousing

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAdditionsWarehousing error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建AdditionsWarehousing对象
func (e *AdditionsWarehousing) Insert(orderId, StoreRoomId int, c *dto.AdditionsWarehousingInsertReq) error {
	var err error
	var data models.AdditionsWarehousing
	c.Generate(&data)
	data.WId = int64(orderId)
	data.StoreRoomId = StoreRoomId
	if c.PurchaseAt != "" {
		if star, err := time.ParseInLocation(time.DateTime, c.PurchaseAt, global.LOC); err == nil {
			data.PurchaseAt = sql.NullTime{
				Time:  star,
				Valid: true,
			}
		}

	} else {
		data.PurchaseAt = sql.NullTime{}
	}

	if c.ExpireAt != "" {
		if end, err := time.ParseInLocation(time.DateTime, c.ExpireAt, global.LOC); err == nil {
			data.ExpireAt = sql.NullTime{
				Time:  end,
				Valid: true,
			}
		}

	} else {
		data.ExpireAt = sql.NullTime{}
	}

	err = e.Orm.Create(&data).Error
	Code := fmt.Sprintf("ZC%08d", data.Id)
	e.Orm.Model(&models.AdditionsWarehousing{}).Where("id = ?", data.Id).Updates(map[string]interface{}{
		"code": Code,
	})
	if err != nil {
		e.Log.Errorf("AdditionsWarehousingService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

func (e *AdditionsWarehousing) Update(uid string, c *dto.AdditionsWarehousingUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AdditionsWarehousing{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, uid)
	c.Generate(&data)

	data.Id = func() int {
		uu, _ := strconv.Atoi(uid)
		return uu
	}()
	if c.PurchaseAt != "" {
		if star, err := time.ParseInLocation(time.DateTime, c.PurchaseAt, global.LOC); err == nil {
			data.PurchaseAt = sql.NullTime{
				Time:  star,
				Valid: true,
			}
		}

	} else {
		data.PurchaseAt = sql.NullTime{}
	}

	if c.ExpireAt != "" {
		if end, err := time.ParseInLocation(time.DateTime, c.ExpireAt, global.LOC); err == nil {
			data.ExpireAt = sql.NullTime{
				Time:  end,
				Valid: true,
			}
		}

	} else {
		data.ExpireAt = sql.NullTime{}
	}
	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AdditionsWarehousingService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

func (e *AdditionsWarehousing) UpdateStore(StoreRoomId int, c *dto.AdditionsWarehousingUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.AdditionsWarehousing{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	data.StoreRoomId = StoreRoomId
	if c.PurchaseAt != "" {
		if star, err := time.ParseInLocation(time.DateTime, c.PurchaseAt, global.LOC); err == nil {
			data.PurchaseAt = sql.NullTime{
				Time:  star,
				Valid: true,
			}
		}

	} else {
		data.PurchaseAt = sql.NullTime{}
	}

	if c.ExpireAt != "" {
		if end, err := time.ParseInLocation(time.DateTime, c.ExpireAt, global.LOC); err == nil {
			data.ExpireAt = sql.NullTime{
				Time:  end,
				Valid: true,
			}
		}

	} else {
		data.ExpireAt = sql.NullTime{}
	}
	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AdditionsWarehousingService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除AdditionsWarehousing
func (e *AdditionsWarehousing) Remove(d *dto.AdditionsWarehousingDeleteReq, p *actions.DataPermission) error {
	var data models.AdditionsWarehousing

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAdditionsWarehousing error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
