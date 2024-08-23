package models

import (
	"database/sql"
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/models"
	"gorm.io/gorm"
)

type RsDial struct {
	models.Model
	Bu               string                 `json:"bu" gorm:"type:varchar(10);comment:bu业务"`
	Desc             string                 `json:"desc" gorm:"type:varchar(35);comment:描述信息"`
	CustomId         int                    `json:"customId" gorm:"type:bigint;comment:所属客户"`
	ContractId       int                    `json:"contractId" gorm:"type:bigint;comment:关联合同"`
	BroadbandType    int                    `json:"broadbandType" gorm:"type:bigint;comment:带宽类型,broadband_type"`
	IsManager        int                    `json:"isManager" gorm:"type:tinyint(1);comment:是否管理线"`
	Account          string                 `json:"account" gorm:"type:varchar(25);comment:账号"`
	Ip               string                 `json:"ip" gorm:"type:varchar(16);comment:IP地址"`
	Pass             string                 `json:"pass" gorm:"type:varchar(20);comment:密码"`
	Mac              string                 `json:"mac" gorm:"type:varchar(30);comment:MAC地址"`
	DialName         string                 `json:"dialName" gorm:"type:varchar(20);comment:线路名称"`
	NetworkingStatus int                    `json:"networkingStatus" gorm:"default:0;type:int;comment:拨号状态 1:已联网 0:待使用 -1:联网异常"`
	Status           int                    `json:"status" gorm:"default:0;type:int;comment:拨号状态,1:已拨通 0:待使用 -1:拨号异常"`
	Source           int                    `json:"source" gorm:"type:int;comment:拨号状态,0:录入 1:自动上报"`
	IdcId            int                    `json:"idcId" gorm:"type:bigint;comment:关联的IDC"`
	HostId           int                    `json:"hostId" gorm:"type:bigint;comment:关联主机ID"`
	DeviceName       string                 `json:"deviceName" gorm:"-"`
	DeviceId         int                    `json:"deviceId" gorm:"type:bigint;comment:关联网卡ID"`
	RunTime          sql.NullTime           `json:"-" gorm:"type:datetime(3);comment:RunTime"`
	RunTimeAt        string                 `json:"runTimeAt" gorm:"-"`
	IspId            int                    `json:"ispId"  gorm:"type:int(1);default:0;comment:关联合同下的账号的运营商ID"`
	IdcInfo          interface{}            `json:"idcInfo" gorm:"-"`
	HostInfo         map[string]interface{} `json:"hostInfo" gorm:"-"`
	models.ModelTime
	models.ControlBy
}

func (RsDial) TableName() string {
	return "rs_dial"
}

func (e *RsDial) AfterFind(tx *gorm.DB) (err error) {

	if e.RunTime.Valid {
		e.RunTimeAt = e.RunTime.Time.Format("2006-01-02 15:04:05")

	}
	e.HostInfo = map[string]interface{}{
		"hostname": "",
		"sn":       "",
	}
	e.IdcInfo = map[string]interface{}{
		"name": "",
	}

	if e.IdcId > 0 {
		var idcRow RsIdc
		tx.Model(&idcRow).Select("id,name").Where("id = ?", e.HostId).Find(&idcRow)
		if idcRow.Id > 0 {
			e.IdcInfo = map[string]interface{}{
				"name": idcRow.Name,
			}
		}
	}
	if e.HostId > 0 {
		var host RsHost
		tx.Model(&host).Select("id,host_name,sn").Where("id = ?", e.HostId).Find(&host)
		if host.Id > 0 {
			e.HostInfo = map[string]interface{}{
				"hostname": host.HostName,
				"sn":       host.Sn,
			}
		}
	}
	if e.DeviceId > 0 {
		var DeviceName models2.HostNetDevice
		tx.Model(&DeviceName).Select("name").Where("id = ?", e.DeviceId).Find(&DeviceName)
		e.DeviceName = DeviceName.Name
	}
	return nil
}
func (e *RsDial) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RsDial) GetId() interface{} {
	return e.Id
}
