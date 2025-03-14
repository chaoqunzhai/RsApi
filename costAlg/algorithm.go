package costAlg

import (
	"errors"
	"fmt"
	"go-admin/cmd/migrate/migration/models"
	"go-admin/common/prometheus"
	"go-admin/common/utils"
	"go-admin/config"
	"gorm.io/gorm"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type CostAlgorithm struct {
	BacktrackInt int
	Orm         *gorm.DB
	Metrics     []*interface{}
	BusinessMap map[int]Business
	RunTime     map[string]string
	PromReq     PromReq
	thiMonthDay int
}

func (c *CostAlgorithm) SetupDb(dbs map[string]*gorm.DB) {
	for _, db := range dbs {
		c.Orm = db
	}
	c.RunTime = make(map[string]string)
}
func (c *CostAlgorithm) BillingMethod() {

}

// 开始获取所有的机器 并且聚合到指定业务下
func (c *CostAlgorithm) StartHostCompute() {
	var BusinessList []models.Business
	hostIds := make([]int, 0)
	nowTime := time.Now()

	c.Orm.Model(&models.Business{}).Find(&BusinessList)
	c.BusinessMap = make(map[int]Business, len(BusinessList))
	for _, bu := range BusinessList {

		var bindHostIds []int
		hostBindBusiness := fmt.Sprintf("SELECT host_id  FROM `host_bind_business` WHERE `host_bind_business`.`business_id` = %v", bu.Id)
		c.Orm.Raw(hostBindBusiness).Scan(&bindHostIds)
		var hostList []models.Host
		//bindHostIdsDemo := []int{1187}
		c.Orm.Model(&models.Host{}).Select("id,host_name,sn,status,balance,idc,isp").Where("id in ? and suspend_billing = 1", bindHostIds).Find(&hostList)
		buValue, ok := c.BusinessMap[bu.Id]
		if !ok {
			buValue.Host = make([]*Host, 0)
		}
		buValue.AlgId = bu.BillingMethod
		buValue.Id = bu.Id
		buValue.Name = bu.Name
		buValue.EnName = bu.EnName
		buValue.IspCnf = make(map[string]*IspCnf)
		buValue.SlaConf = map[int]*SlaConf{
			1: {
				Id:    1,
				Name:  "offLineRange",
				Start: "18:00",
				End:   "24:00",
				Limit: 30,
				Desc:  "当日机器晚高峰期间线路断开超30min，该线路流量当日不计费",
			},
		}

		for _, host := range hostList {
			//获取每 host -> IDC  -> 客户 -> 合同 -> 费用
			HostRow :=&Host{
				HostId:          host.Id,
				IspId:           host.Isp,
				HostName:        host.HostName,
				HostSn:          host.Sn,
				IdcId:           host.Idc,
				BandwidthIncome: host.Balance / 1000, //换算成G
				Balance: host.Balance,
			}
			note:=make([]string,0)
			if host.Idc > 0 {

				var idcRow models.Idc
				c.Orm.Model(&models.Idc{}).Select("custom_id,name,id").Where("id = ?",host.Idc).Limit(1).Find(&idcRow)

				if idcRow.CustomId > 0 {


					var CustomRow models.Custom
					c.Orm.Model(&models.Custom{}).Select("name,id").Where("id = ?",idcRow.CustomId).Limit(1).Find(&CustomRow)

					if CustomRow.Id > 0 {

						var ContractRow models.Contract
						c.Orm.Model(&models.Contract{}).Select("id,name").Where("custom_id = ?", idcRow.CustomId).Limit(1).Find(&ContractRow)
						if ContractRow.Id > 0 {
							var BandwidthFeesRow models.BandwidthFees
							c.Orm.Model(&models.BandwidthFees{}).Where("contract_id = ? and isp = ?", ContractRow.Id,host.Isp).Limit(1).Find(&BandwidthFeesRow)

							if BandwidthFeesRow.Id > 0 {
								//最后匹配到了 才会设置对象
								HostRow.ContractAlg = &ContractAlg{
									IspId: BandwidthFeesRow.Isp,
									LinePrice:BandwidthFeesRow.LinePrice,
									ManagerLineCost: BandwidthFeesRow.ManagerLineCost,
								}
							}else {
								note  = append(note,fmt.Sprintf("合同:%v-%v未关联费用信息",ContractRow.Id,ContractRow.Name ))
							}

						}else {
							note  = append(note,fmt.Sprintf("客户:%v-%v未关联合同",CustomRow.Id,CustomRow.Name ))
						}


					}else {
						note  = append(note,fmt.Sprintf("IDC:%v-%v未关联客户",idcRow.Id,idcRow.Name))
					}

				}else {
					note  = append(note,fmt.Sprintf("IDC:%v-%v未关联客户",idcRow.Id,idcRow.Name))
				}
			}else {
				note  = append(note,"未关联IDC")
			}
			if len(note) > 0 {
				HostRow.AlgNote = note
			}
			buValue.Host = append(buValue.Host, HostRow)
			hostIds = append(hostIds, host.Id)
		}

		c.BusinessMap[bu.Id] = buValue
	}

	endTime := time.Now()
	hostIds = utils.RemoveRepeatInt(hostIds)
	c.RunTime["1.StartHostCompute"] = fmt.Sprintf("获取所有业务下机器耗时:%v", endTime.Sub(nowTime))

	//构造价格
	c.BuPrice()
	//构造主机的业务SN
	c.HostBuSn(hostIds)
	//开始计算
	c.BacktrackComputeMixedAlg()
}
func (c *CostAlgorithm) HostBuSn(hostIds []int) {
	nowTime := time.Now()
	businessSnList := make([]models.HostSoftware, 0)
	c.Orm.Model(&models.HostSoftware{}).Where("host_id in ? and `key` LIKE 'sn_%'", hostIds).Find(&businessSnList)
	hostMap := make(map[int][]*LabelRow)

	for _, item := range businessSnList {
		if strings.HasPrefix(item.Key, "sn_") {
			itemKey := strings.Replace(item.Key, "sn_", "", -1)

			snLabel := &LabelRow{
				Label: itemKey,
				Value: item.Value,
			}
			snList, ok := hostMap[item.HostId]
			if !ok {
				snList = make([]*LabelRow, 0)
			}
			snList = append(snList, snLabel)

			hostMap[item.HostId] = snList
		}
	}
	endTime := time.Now()

	for _, bu := range c.BusinessMap {
		for _, host := range bu.Host {

			if snList, ok := hostMap[host.HostId]; ok {
				host.BuSn = snList
			}
		}
		c.BusinessMap[bu.Id] = bu
	}
	c.RunTime["2.HostBuSn"] = fmt.Sprintf("获取主机的业务SN,和组装耗时:%v", endTime.Sub(nowTime))

}

// 获取不同 业务 - 运营商的计费价格
func (c *CostAlgorithm) BuPrice() {
	nowTime := time.Now()

	c.thiMonthDay = c.GetMonthDay()
	for _, bu := range c.BusinessMap {

		var BusinessCostCnf []models.BusinessCostCnf

		c.Orm.Model(&BusinessCostCnf).Where("bu_id = ?", bu.Id).Find(&BusinessCostCnf)

		for _, cnf := range BusinessCostCnf {
			var ispStr string
			switch cnf.Isp {
			case 1:
				ispStr = "移动"
			case 2:
				ispStr = "电信"
			case 3:
				ispStr = "联通"
			default:
				ispStr = "其他"
			}

			bu.IspCnf[ispStr] = &IspCnf{
				ConstId:     cnf.Isp,
				Name:        ispStr,
				Price:       cnf.Price,
				Day:         c.thiMonthDay,
				AvgBuDayPrice: utils.RoundDecimalFlot64(3, cnf.Price/float64(c.thiMonthDay)),
			}
		}

	}

	endTime := time.Now()

	c.RunTime["2.RichBu"] = fmt.Sprintf("丰富业务的价格配置耗时:%v", endTime.Sub(nowTime))

}

func (c *CostAlgorithm) BacktrackComputeMixedAlg() {
	for i := 0; i < c.BacktrackInt; i++ {
		now := time.Now()
		// 计算昨天的日期
		yesterday := now.AddDate(0, 0, -i)

		// 计算昨天的 0 点
		startOfYesterday := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, yesterday.Location())

		// 计算昨天的 23:59:59
		endOfYesterday := startOfYesterday.Add(24*time.Hour - time.Second)

		c.ComputeMixedAlg(startOfYesterday,endOfYesterday)
	}
}
// 获取即可 进行费用计算 + 利用率计算 + SLA计算
func (c *CostAlgorithm) ComputeMixedAlg(startOfYesterday,endOfYesterday time.Time) {

	now := time.Now()
	c.PromReq.Setup = 300
	c.PromReq.Start = fmt.Sprintf("%v", startOfYesterday.Unix())
	c.PromReq.End = fmt.Sprintf("%v", endOfYesterday.Unix())

	for _, bu := range c.BusinessMap {

		for _, host := range bu.Host {

			data := c.GetHostPrometheusData(host, &bu)
			host.AlgDay = startOfYesterday.Format(time.DateOnly)

			host.PriceCompute = data

			host.BuId = bu.Id

			//fmt.Printf("host:%+v,GetHostPrometheusData:%+v\n", host, data)
			c.InsertDb(host)
		}
		c.BusinessMap[bu.Id] = bu
	}
	endTime := time.Now()
	c.RunTime["2.RichBu"] = fmt.Sprintf("计算所有主机PrometheusData数据耗时:%v", endTime.Sub(now))

}

func (c *CostAlgorithm) GetHostPrometheusData(host *Host, bu *Business) map[string]*MonitorCompute {
	//获取昨天日期的 开始和结束

	//请求Prometheus进行日95计算数据

	algMap := make(map[string]*MonitorCompute)
	for _, isp := range bu.IspCnf {
		if host.IspId != isp.ConstId {
			continue
		}
		var transmitQuery string
		isMb := false
		switch bu.EnName {
		case "dianxin":
			isMb = true
			transmitQuery = fmt.Sprintf("sum(flow_bandwidth_by_minute{instance=\"%v\"})", host.HostName)
		default:
			transmitQuery = fmt.Sprintf("sum(rate(phy_nic_network_transmit_bytes_total{instance=\"%v\",device_isp=\"%v\"}[5m]))*8", host.HostName, isp.Name)

		}

		compute := c.requestPromResult(isMb, transmitQuery, bu.SlaConf, isp)

		if compute.Empty { //如果是空数据 在其他指标中进行请求
			continue
		}
		//利用率 =  95 / 总带宽
		usage := utils.RoundDecimalFlot64(3, compute.PercentG/host.BandwidthIncome)
		if usage > 1 {
			usage = 1
		}
		compute.Usage = usage
		algMap[isp.Name] = compute
	}

	return algMap
}

func (c *CostAlgorithm) requestPromResult(isMb bool, query string, SlaConf map[int]*SlaConf, IspCnf *IspCnf) *MonitorCompute {

	result := &MonitorCompute{}
	//查询普罗米修斯数据
	queryUrl, err := url.Parse(func() string {
		vv, _ := url.JoinPath(config.ExtConfig.Prometheus.Endpoint, "/api/v1/query_range")
		return vv
	}())

	parameters := url.Values{}
	parameters.Add("start", c.PromReq.Start)
	parameters.Add("end", c.PromReq.End)
	parameters.Add("step", fmt.Sprintf("%v", c.PromReq.Setup))
	parameters.Add("query", query)
	queryUrl.RawQuery = parameters.Encode()

	ProResult, err := prometheus.GetPromRangeResult(queryUrl)

	if err != nil {
		return result
	}
	XData := make([][]interface{}, 0)
	XValue := make([]float64, 0)

	if len(ProResult.Data.Result) > 0 {

		parentUnix := 0.0 //上一个时间戳

		setupNumber := 0.0 //时间戳的区间
		for _, row := range ProResult.Data.Result[0].Value {

			if len(row) != 2 {
				continue
			}
			unixFloat := row[0].(float64)
			if setupNumber == 0 { //只需要计算一次区间即可
				setupNumber = unixFloat - parentUnix //计算区间
			}
			if parentUnix > 0 { //如果有上一个值
				//时间戳的差大于区间,那就少点了,补多少点,那就算一个次数循环写入
				ac := unixFloat - parentUnix
				if ac > setupNumber {
					for i := 0; i < int(ac/setupNumber); i++ {
						addUnix := unixFloat + setupNumber*float64(i+1)

						addUnixTime := time.Unix(int64(addUnix), 0)
						//进行计算
						XData = append(XData, []interface{}{addUnixTime, 0})
						XValue = append(XValue, 0)

					}
				}

			}
			parentUnix = unixFloat

			valueStr := row[1].(string)
			valueFloat, _ := strconv.ParseFloat(valueStr, 64)
			unixTime := time.Unix(int64(unixFloat), 0)
			XData = append(XData, []interface{}{unixTime, valueStr})

			//进行计算
			XValue = append(XValue, valueFloat)

		}
	}
	sort.Float64s(XValue)

	//计算95值

	if len(XValue) > 1 {

		Percent := utils.Percentile(XValue, 0.95)
		result.PercentBytes = Percent

		if isMb { //是MB的单位
			result.PercentG = utils.RoundDecimalFlot64(3, Percent/1000)
		} else {
			result.PercentG = utils.RoundDecimalFlot64(3, Percent/(1024*1024*1024))
		}

		//计算max
		result.Max = utils.Max(XValue)
		//计算最小
		result.Min = utils.Min(XValue)

		//

		result.HeartbeatNum = len(XValue)

		result.TotalBandwidth = utils.SumFloats(XValue)
		//计算平均
		result.Avg = utils.Avg(XValue)
		//SLA计算
		result.SLA = c.AlgSla(XData, SlaConf)

		result.IspDayPrice = utils.RoundDecimalFlot64(3, IspCnf.AvgBuDayPrice*result.PercentG)
		result.IspCnf = IspCnf
	} else {
		result.Empty = true
	}

	return result
}

// 对于混跑的业务,是有比例的,
func (c *CostAlgorithm) Mixed() {
	//例如: 白山-金山的机器 ,那主要就是 跑白山的，白山一天跑200个G 那金山只跑5%

}

// 触发SLA的算法
func (c *CostAlgorithm) AlgSla(data [][]interface{}, SlaConf map[int]*SlaConf) (sla *SlaRow) {
	//当日出现1次机器故障，当日不计费([高峰期]当日10-14点超过15min或当日18~24时设备超过5min末服务视为机器改南
	//当日机器晚高峰期间线路断开超30min，该线路流量当日不计费
	sla = &SlaRow{}

	for _, cnf := range SlaConf {
		//多个规格依次匹配,如果匹配到了一个规则就跳出
		limitTag := 0
		for _, entry := range data {
			var startTime time.Time
			var err error
			if startTime, err = ParseTime(cnf.Start); err != nil {
				continue
			}
			var endTime time.Time
			if endTime, err = ParseTime(cnf.End); err != nil {
				continue
			}
			timestamp := time.Unix(int64(entry[0].(float64)), 0)
			if timestamp.After(startTime) && timestamp.Before(endTime) {
				limitTag += 5
			}
		}

		if limitTag >= 30 {
			return &SlaRow{
				Trigger: false,
				Info:    cnf.Desc,
			}
		}
	}
	return
}

// 日95 计算
func (c *CostAlgorithm) Day95() {

}

// 月95
func (c *CostAlgorithm) Month95() {

}

// [day95相加] / 月天数
func (c *CostAlgorithm) MonthAvgDay95() {

}

// 单个运营商下的 [日95相加]  / 月天数
func (c *CostAlgorithm) IspMonthAvgDay95() {}

// 每天的日峰值相加 / 月天数
func (c *CostAlgorithm) DayUpMonthAvg() {}

// 买断
func (c *CostAlgorithm) Buyout() {}

// 日95月平均阶梯计费
func (c *CostAlgorithm) RangeMonthAvgDay95() {}

// 当月的自然月天数
func (c *CostAlgorithm) GetMonthDay() int {
	// 获取当前时间
	now := time.Now()

	// 获取当前年份和月份
	year, month, _ := now.Date()

	// 获取当前月份的第一天
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())

	// 获取下个月的第一天
	nextMonthFirst := firstDay.AddDate(0, 1, 0)

	// 计算当前月份的天数（下个月第一天减去当前月份第一天，然后减去时间间隔中的一天）
	daysInMonth := int(nextMonthFirst.Sub(firstDay).Hours() / 24)

	return daysInMonth
}
func (c *CostAlgorithm)getYesterdayInCurrentMonth() (string, error) {
	// 获取当前时间
	now := time.Now()

	// 计算昨天的日期
	yesterday := now.AddDate(0, 0, -1)

	// 检查昨天是否还在当前月
	if yesterday.Month() != now.Month() {
		return "", errors.New("昨天是上个月的日期")
	}

	// 返回昨天的日期，格式为 YYYY-MM-DD
	return yesterday.Format("2006-01-02"), nil
}
func (c *CostAlgorithm) InsertDb(host *Host) {

	//需要拆分字段
	for _, row := range host.PriceCompute { //不同的运营商的计费

		var HostIncome models.HostIncome

		c.Orm.Model(&HostIncome).Where("host_id = ? and alg_day = ?", host.HostId, host.AlgDay).Limit(1).Find(&HostIncome)


		//做一次成本价的计算
		//合同宽费用 单条多少钱 *  这个机器是几条线 / 自然月  =  机器每天的成本

		//因为可能不同的商务理解的不一样， 如果是>20 那就是写的 价格/30 就是每天每M
		var (
			DayCost float64
			CostAlgorithmVal string
		)
		//因为还涉及到了 一个停止计费, 那当月成本就是 算的每天的成本 然后叠加起来 才是当月的成本
		if host.ContractAlg  !=nil{
			var mbPrice float64
			if host.ContractAlg.LinePrice > 20 {
				mbPrice = host.ContractAlg.LinePrice  / float64(utils.GetDaysInMonth(time.Now().Year(),time.Now().Month()))
			}else {
				mbPrice = host.ContractAlg.LinePrice
			}

			DayCost = (host.Balance * mbPrice) / float64(c.thiMonthDay)
			CostAlgorithmVal = fmt.Sprintf("(合同单价(%v)*主机总带宽(%v)) / %v = 每天成本,月成本=每天计算成本相加",mbPrice,host.Balance,c.thiMonthDay)
		}else {
			CostAlgorithmVal = strings.Join(host.AlgNote,",")
		}
		//否则就是 业务线单价 元/M/月
		//var yesterDayFloat float64
		//if yesterDay,err:=c.getYesterdayInCurrentMonth();err!=nil{
		//	var oldDayHostIncome models.HostIncome
		//
		//	c.Orm.Model(&oldDayHostIncome).Select("id,day_cost").Where("host_id = ? and alg_day = ?", host.HostId,yesterDay).Limit(1).Find(&oldDayHostIncome)
		//	if oldDayHostIncome.Id > 0 {
		//		yesterDayFloat = oldDayHostIncome.DayCost
		//	}
		//
		//}
		DayCost =  utils.RoundDecimalFlot64(4,DayCost)
		if HostIncome.Id > 0 {
			HostIncome.Isp = row.IspCnf.ConstId
			HostIncome.IdcId = host.IdcId
			HostIncome.BuId = host.BuId
			HostIncome.DayCost = DayCost
			HostIncome.Income = row.IspDayPrice
			HostIncome.Usage = row.Usage
			HostIncome.Bandwidth95 = row.PercentG
			HostIncome.SlaInfo = row.SLA.Info
			HostIncome.SlaPrice = row.SLA.Price
			HostIncome.HeartbeatNum = row.HeartbeatNum
			HostIncome.TotalBandwidth = row.TotalBandwidth
			HostIncome.AvgBuDayPrice = row.IspCnf.AvgBuDayPrice
			HostIncome.HostId = host.HostId
			HostIncome.AlgDay = host.AlgDay
			HostIncome.CostAlgorithm  = CostAlgorithmVal
			c.Orm.Save(&HostIncome)
			continue
		}
		RsHostIncome := models.HostIncome{
			AlgDay:         host.AlgDay,
			HostId:         host.HostId,
			DayCost:DayCost,
			Isp:            row.IspCnf.ConstId,
			IdcId:          host.IdcId,
			BuId:           host.BuId,
			Income:         row.IspDayPrice,
			Usage:          row.Usage,
			Bandwidth95:    row.PercentG,
			SlaInfo:        row.SLA.Info,
			SlaPrice:       row.SLA.Price,
			HeartbeatNum:   row.HeartbeatNum,
			CostAlgorithm: CostAlgorithmVal,
			TotalBandwidth: row.TotalBandwidth,
			AvgBuDayPrice:    row.IspCnf.AvgBuDayPrice,
		}
		c.Orm.Create(&RsHostIncome)
	}
}

