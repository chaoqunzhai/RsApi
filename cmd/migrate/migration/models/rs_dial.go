package models

import "database/sql"

type Dial struct {
	RichGlobal
	CustomId         int          `json:"customId" gorm:"comment:所属客户"`
	ContractId       int          `json:"contractId" gorm:"comment:关联合同"`
	BroadbandType    int          `json:"broadbandType" gorm:"comment:带宽类型,broadband_type"`
	IsManager        int64        `json:"isManager" gorm:"type:tinyint(1);comment:是否管理线"`
	Account          string       `json:"account" gorm:"type:varchar(25);comment:账号"`
	Ip               string       `json:"ip"  gorm:"type:varchar(16);comment:IP地址"`
	Pass             string       `json:"pass" gorm:"type:varchar(20);comment:密码"`
	Mac              string       `json:"mac" gorm:"type:varchar(30);comment:MAC地址"`
	DialName         string       `json:"dialName" gorm:"type:varchar(20);comment:线路名称"`
	NetworkingStatus int          `json:"networkingStatus" gorm:"type:int(1);default:0;comment:拨号状态,1:已联网 0:未联网 -1:联网异常"`
	Status           int          `json:"status" gorm:"type:int(1);default:0;comment:拨号状态,1:已拨通 0:待使用 -1:拨号异常"`
	Source           int          `json:"source" gorm:"type:int(1);default:0;comment:拨号状态,0:录入 1:自动上报"`
	IspId            int          `json:"ispId"  gorm:"type:int(1);default:0;comment:关联合同下的账号的运营商ID"`
	IdcId            int          `json:"idcId" gorm:"index;comment:关联的IDC"`
	HostId           int          `json:"hostId" gorm:"index;comment:关联主机ID"`
	DeviceId         int          `json:"deviceId" gorm:"index;comment:关联网卡ID"`
	RunTime          sql.NullTime `json:"runTime"`
}

func (Dial) TableName() string {
	return "rs_dial"
}
