/*
*
@Author: chaoqun
* @Date: 2024/7/25 22:35
*/
package apis

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/cmdb/models"
	"go-admin/app/cmdb/service/dto"
	models2 "go-admin/cmd/migrate/migration/models"
	_ "go-admin/common/dial"
	models3 "go-admin/common/models"
	"go-admin/global"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type RegisterApi struct {
	api.Api
}

var ispList = []string{
	"电信",
	"联通",
	"移动",
}

// 黑名单的SN， 因为有些SN都是一样的，只能通过主机名来确定唯一性
var blackMap = map[string]bool{
	"01234567890123456789AB": true,
	"Default string":         true,
}

func RemoveBracketContent(s string) string {
	var builder strings.Builder
	skip := false
	for _, r := range s {
		if r == '(' {
			skip = true
		} else if r == ')' {
			skip = false
		} else if !skip {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
func (e *RegisterApi) InitIdc(req dto.RegisterMetrics) int {

	re := regexp.MustCompile(`\d+`)

	matches := re.FindStringSubmatch(req.Remark[0:8])
	var err error
	var idcStrNumberIndex int
	var idcNumber int
	if len(matches) == 0 {
		return 0
	}
	idcStrNumber := matches[0]

	idcNumber, err = strconv.Atoi(idcStrNumber)
	if err != nil {
		return 0
	}
	idcStrNumberIndex = len(idcStrNumber)

	removeNumberStr := req.Remark[idcStrNumberIndex:]

	idcName := ""
	for _, k := range ispList {
		ispIndex := strings.Index(removeNumberStr, k)
		if ispIndex < 0 {
			continue
		}
		idcName = removeNumberStr[:ispIndex]
	}
	if idcName == "" {
		idcName = RemoveBracketContent(removeNumberStr)
	}

	var IdcRow models.RsIdc

	e.Orm.Model(&models.RsIdc{}).Select("id").Where("number = ?", idcNumber).Limit(1).Find(&IdcRow)

	if IdcRow.Id == 0 { //进行创建IDC

		var regionId []string
		if req.Province != "" {
			var Province models2.ChinaData
			e.Orm.Model(&models2.ChinaData{}).Select("id").Where("name like ?", "%"+req.Province+"%").Limit(1).Find(&Province)
			if Province.Id > 0 {
				regionId = append(regionId, fmt.Sprintf("%v", Province.Id))
			}

		}
		if req.City != "" {
			var City models2.ChinaData
			e.Orm.Model(&models2.ChinaData{}).Select("id").Where("name like ?", "%"+req.City+"%").Limit(1).Find(&City)
			if City.Id > 0 {
				regionId = append(regionId, fmt.Sprintf("%v", City.Id))
			}
		}

		IdcRow = models.RsIdc{
			Number: idcNumber,
			Name:   idcName,
			Desc:   req.Remark,
			Status: 1,
			Belong: 1,
			Region: strings.Join(regionId, ","),
		}
		e.Orm.Create(&IdcRow)
	}
	return IdcRow.Id
}

// RegisterData
// @Summary 主机存活注册
// @Description 主动上报
// @Tags 主机上报
// @Success 200 {object} response.Response "{"code": 200, "data": "","msg":"successful"}"
// @Router /api/v1/register/healthy [post]

func (e *RegisterApi) Healthy(c *gin.Context) {

	req := dto.RegisterMetrics{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		Errors
	if err != nil {
		e.Logger.Error(err)
		body, readErr := ioutil.ReadAll(c.Request.Body)
		if readErr != nil {
			e.Error(500, err, err.Error())
			return
		}
		defer func() {
			_ = c.Request.Body.Close()
		}()
		e.Logger.Error(fmt.Sprintf("post Body: %v", string(body)))
		e.Error(500, err, err.Error())
		return
	}
	registerHeaderKey := c.GetHeader("RsRole")

	if strings.TrimSpace(registerHeaderKey) != "rs-sre" {

		e.Error(http.StatusUnauthorized, nil, "you set Header")
		return
	}

	var hostInstance models.RsHost

	SN := strings.TrimSpace(req.Sn)
	HOSTNAME := strings.TrimSpace(req.Hostname)
	//SN是否为一个 黑名单.如果是 用主机名做唯一性校验
	isDirty := blackMap[SN]
	if isDirty {
		e.Orm.Model(&hostInstance).Where("host_name = ?", HOSTNAME).First(&hostInstance)
	} else {
		e.Orm.Model(&hostInstance).Where("sn = ?", SN).First(&hostInstance)
	}
	if req.Belong == 0 { //如果为空,那也算是一个自建的机器
		req.Belong = 1
	}
	hostInstance.Belong = req.Belong
	hostInstance.Sn = SN
	hostInstance.HostName = HOSTNAME
	//hostInstance.Ip = req.Ip

	if req.NetType > 0 { //只有非0是才会进行保存
		hostInstance.NetworkType = req.NetType
	}
	hostInstance.Mac = req.Mac
	hostInstance.Mask = req.Mask
	hostInstance.Gateway = req.Gateway
	//hostInstance.PublicIp = req.PublicIp
	hostInstance.Cpu = req.CPU
	//hostInstance.Kernel = req.Kernel
	hostInstance.RemotePort = req.RemotePort
	hostInstance.Status = global.HostSuccess
	hostInstance.Memory = req.Memory
	hostInstance.Remark = req.Remark

	var ispNumber int
	switch strings.TrimSpace(req.Isp) {
	case "移动":
		ispNumber = 1
	case "电信":
		ispNumber = 2
	case "联通":
		ispNumber = 3
	default:
		ispNumber = 4
	}
	if req.Balance > 0 {
		hostInstance.Balance = req.Balance
	}
	if req.BandwidthCnf.Line > 0 {
		hostInstance.AllLine = int(req.BandwidthCnf.Line)
		hostInstance.LineBandwidth = req.BandwidthCnf.Width
	}
	hostInstance.Isp = ispNumber

	hostInstance.HealthyAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	//Business 如果不为空,进行关联
	if req.Business != "" {
		bindBuList := make([]models.RsBusiness, 0)
		for _, row := range strings.Split(req.Business, "-") {

			var buInstance models.RsBusiness
			e.Orm.Model(&models.RsBusiness{}).Where("en_name = ?", strings.TrimSpace(row)).Limit(1).Find(&buInstance)
			if buInstance.Id > 0 {
				bindBuList = append(bindBuList, buInstance)
			}
		}
		if len(bindBuList) > 0 {
			hostInstance.Business = bindBuList
		}
	}
	// 关联机房 例如解析remark=166陕西延安宜川集义郭东机房电信1-1-10(40*100M)  大概截取前10个字符，考虑到后期可能机房数达上万个
	// 备注不为空 并且 没有关联IDC,那就主动关联

	if req.Remark != "" && len(req.Remark) >= 8 {
		if idcId := e.InitIdc(req); idcId > 0 {
			hostInstance.Idc = idcId
		}

	}

	e.Orm.Save(&hostInstance)

	NetDeviceMap := make(map[string]int, 0)
	if req.NetDevice != "" {

		NetDeviceList := strings.Split(req.NetDevice, ",")

		for _, NetDeviceName := range NetDeviceList {
			var (
				DeviceRow models2.HostNetDevice
			)
			NetDeviceInfo := strings.Split(NetDeviceName, ":")
			var status int
			var DeviceName string
			if len(NetDeviceInfo) > 1 { //有状态
				DeviceName = NetDeviceInfo[0]
				DeviceStatus := NetDeviceInfo[1]
				if DeviceStatus == "UP" {
					status = 1
				} else {
					status = -1
				}
			} else {
				DeviceName = NetDeviceName
				status = 1
			}
			if strings.HasSuffix(DeviceName, "UNKNOWN") { //没有识别的网卡不做处理
				continue
			}
			if strings.HasPrefix(DeviceName, "lo") { //回环端口
				continue
			}
			if strings.HasPrefix(DeviceName, "docker") {
				continue
			}
			if strings.HasPrefix(DeviceName, "ppp") {
				continue
			}
			//获取的网卡有的是@拼接的
			DeviceNameList := strings.Split(DeviceName, "@")
			if len(DeviceNameList) == 0 {
				continue
			}
			inDbDeviceName := DeviceNameList[0]
			e.Orm.Model(&models2.HostNetDevice{}).Where("host_id = ? and `name` = ?",
				hostInstance.Id, inDbDeviceName).First(&DeviceRow)
			//如果 DeviceName 有 @ 那就需要特殊处理

			DeviceRow.HostId = hostInstance.Id
			DeviceRow.Name = inDbDeviceName
			DeviceRow.UpdatedAt = models3.XTime{
				Time: time.Now(),
			}
			DeviceRow.Status = status
			e.Orm.Save(&DeviceRow)
			NetDeviceMap[inDbDeviceName] = DeviceRow.Id
		}

	}
	//拨号列表,
	for _, DialRow := range req.Dial {

		var (
			DialRowModel models.RsDial
			DialCount    int64
		)
		var bindNetDeviceId int
		NetDeviceId, NetDeviceOk := NetDeviceMap[DialRow.I]
		if NetDeviceOk {
			bindNetDeviceId = NetDeviceId
		} else {
			//没有那就创建一个
			HostNetDevice := models2.HostNetDevice{
				HostId: hostInstance.Id,
				Name:   DialRow.I,
				Status: 1,
			}
			e.Orm.Create(&HostNetDevice)
			bindNetDeviceId = HostNetDevice.Id
		}
		//对于自动上报数据的数据,做一个特定创建,防止 已经创建了这个账号，被自动创建也冲掉
		e.Orm.Model(&models.RsDial{}).Where("account = ? and source = 1", DialRow.A).Count(&DialCount)
		// DialRow.S === 1 已经拨通了,但是是否可以上网 还需要进行检测
		// DialRow.S === 1 拨号失败了,那联网也是失败的

		if DialCount > 0 {
			updateMap := map[string]interface{}{
				"host_id":   hostInstance.Id,
				"idc_id":    hostInstance.Idc,
				"account":   DialRow.A,
				"pass":      DialRow.P,
				"status":    DialRow.S,
				"ip":        DialRow.Ip,
				"mac":       DialRow.Mac,
				"source":    1,
				"dial_name": DialRow.D,
				"bu":        DialRow.BU,
				"ip_v6":     DialRow.IpV6,
				"nat_type":  DialRow.NT,
				"vlan_id":   DialRow.VlanId,
				"device_id": bindNetDeviceId,
			}
			if DialRow.NS != 0 {
				updateMap["networking_status"] = DialRow.NS
			}
			e.Orm.Model(&models.RsDial{}).Where("account = ? and source = 1", DialRow.A).Updates(updateMap)
		} else {
			DialRowModel.Bu = DialRow.BU
			DialRowModel.HostId = hostInstance.Id
			DialRowModel.IdcId = hostInstance.Idc
			DialRowModel.Account = DialRow.A
			DialRowModel.Pass = DialRow.P
			DialRowModel.Ip = DialRow.Ip
			DialRowModel.IpV6 = DialRow.IpV6
			DialRowModel.VlanId = DialRow.VlanId
			DialRowModel.Mac = DialRow.Mac
			DialRowModel.DialName = DialRow.D
			DialRowModel.NatType = DialRow.NT
			DialRowModel.Source = 1
			DialRowModel.DeviceId = bindNetDeviceId
			DialRowModel.Status = DialRow.S

			if DialRow.NS != 0 {
				DialRowModel.NetworkingStatus = DialRow.NS
			}
			e.Orm.Save(&DialRowModel)
		}
	}
	var hostSystem models2.HostSystem

	e.Orm.Model(&models2.HostSystem{}).Where("host_id = ?", hostInstance.Id).First(&hostSystem)
	MemoryData := func() string {

		dat, _ := json.Marshal(req.MemoryMap)

		return string(dat)
	}()
	hostSystem.Disk = req.Disk
	hostSystem.TransmitNumber = req.TransmitNumber
	hostSystem.ReceiveNumber = req.ReceiveNumber
	hostSystem.HostId = hostInstance.Id
	hostSystem.MemoryData = MemoryData
	e.Orm.Save(&hostSystem)

	if len(req.BusinessSn) > 0 {
		for key, val := range req.BusinessSn {
			var (
				snRow models2.HostSoftware
			)
			if val == "" {
				continue
			}
			snKey := fmt.Sprintf("sn_%v", key)
			e.Orm.Model(&models2.HostSoftware{}).Where("host_id = ? and `key` = ?",
				hostInstance.Id, snKey).First(&snRow)
			snRow.HostId = hostInstance.Id
			snRow.Key = snKey
			snRow.Value = val
			e.Orm.Save(&snRow)
		}
	}
	for _, softWare := range req.ExtendMap {
		var (
			softRow models2.HostSoftware
		)
		e.Orm.Model(&models2.HostSoftware{}).Where("host_id = ? and `key`  = ?",
			hostInstance.Id, softWare.Key).First(&softRow)

		softRow.HostId = hostInstance.Id
		softRow.Key = softWare.Key
		softRow.Value = softWare.Value
		softRow.Desc = softWare.Desc
		e.Orm.Save(&softRow)

	}
	e.OK("", "successful")
	return
}
