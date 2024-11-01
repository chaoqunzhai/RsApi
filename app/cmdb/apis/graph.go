package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service/dto"
	"gorm.io/gorm"
	"strings"
)

type Graph struct {
	api.Api
}

func (e Graph) MakeOrm(req *dto.GraphPageReq, orm *gorm.DB) *gorm.DB {

	if req.IdcId != "" {
		orm = orm.Where("idc_id = ? ", req.IdcId)
	}

	if req.HostId != "" {
		orm = orm.Where("host_id = ? ", req.HostId)
	}

	if req.StartTime != "" && req.EndTime != "" {
		orm = orm.Where("created_at >= ? and created_at < ?", req.StartTime, req.EndTime)
	}

	if req.BuId != "" {
		orm = orm.Where("bu_id in ? ", strings.Split(req.BuId, ","))
	}
	if req.IspId != "" {
		orm = orm.Where("isp = ? ", req.IspId)
	}
	if req.CustomId != "" {
		//先通过客户搜索机房
		//搜索到的机房 在收益中查询
		var idcList []string
		e.Orm.Model(&models.RsIdc{}).Where("custom_id = ?", req.CustomId).Find(&idcList).Scan(&idcList)
		orm = orm.Where("idc_id in ?", idcList)
	}
	if req.Region != "" {
		RegionList := strings.Split(req.Region, ",")
		var searchRegion string

		if len(RegionList) > 1 {

			searchRegion = RegionList[len(RegionList)-1]
		} else {
			searchRegion = req.Region
		}
		likeQ := fmt.Sprintf("region like '%%%s%%'", searchRegion)

		var idcList []int
		e.Orm.Model(&models.RsIdc{}).Select("id").Where(likeQ).Find(&idcList).Scan(&idcList)
		orm = orm.Where("idc_id in ?", idcList)
	}
	return orm
}

func (e Graph) Income(c *gin.Context) {
	req := dto.GraphPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var model models.RsHostIncome
	orm := e.Orm.Model(&model)
	orm = e.MakeOrm(&req, orm)

	var buList []models.RsBusiness
	e.Orm.Model(&models.RsBusiness{}).Select("id,name").Find(&buList)

	buMap := make(map[int]string)
	for _, b := range buList {
		buMap[b.Id] = b.Name
	}

	var list []models.RsHostIncome
	orm.Find(list)
	//[
	//	{ 访问时间: '2024-1-12', 业务1: 320, 业务2: 120, 业务3: 220, 业务4: 1500 },
	//	{ 访问时间: '2024-1-13', 业务1: 320, 业务2: 120, 业务3: 220, 业务4: 1500 },
	//	{ 访问时间: '2024-1-14', 业务1: 320, 业务2: 120, 业务3: 220, 业务4: 1500 },
	//	]

	var result []interface{}

	dayMap := make(map[string]map[string]float64, 0) //存放每天业务的数据
	for _, row := range list {

		data, ok := dayMap[row.AlgDay]
		if !ok {
			data = make(map[string]float64, 0)
		}
		buName := buMap[row.BuId]
		data[buName] += row.Income

		dayMap[row.AlgDay] = data
	}

	for k, v := range dayMap {
		row := map[string]interface{}{
			"day": k,
		}
		for k2, v2 := range v {
			row[k2] = v2
		}
		result = append(result, row)
	}
	e.OK(result, "successful")
	return
}
