package api

import (
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/cmd/migrate/migration/models"
	"go-admin/common/dial"
)

func InitializationDialMap() {
	dbList := sdk.Runtime.GetDb()

	//fmt.Println("加载全部的拨号列表")
	for _, d := range dbList {
		var dialList []models.Dial
		d.Model(&models.Dial{}).Select("account,idc_id").Where("enable = ?", true).Find(&dialList)
		for _, row := range dialList {

			event := dial.MapEvent{
				Idc: row.IdcId,
			}
			dial.MapCnf.Set(row.Account, &event)
		}
	}
}
