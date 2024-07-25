package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/cmdb/service/dto"
	models2 "go-admin/cmd/migrate/migration/models"
)

type PublicApi struct {
	service.Service
}

func (e *PublicApi) GetTreeData(pid int, cache []interface{}, lists []models2.ChinaData) (data []interface{}, hs bool) {
	filterList := make([]models2.ChinaData, 0)

	for _, row := range lists {
		if row.Pid == pid {
			filterList = append(filterList, row)
		}
	}
	if len(filterList) == 0 {

		return cache, false
	}

	for _, dat := range filterList {

		row := dto.CityTreeRow{
			Code: dat.Id,
			Name: dat.Name,
		}
		cacheChildren := make([]interface{}, 0)
		Children, valid := e.GetTreeData(dat.Id, cacheChildren, lists)
		if valid {
			row.Children = Children
		}
		cache = append(cache, row)
	}

	return cache, true
}
