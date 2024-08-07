package models

import "database/sql"

type Dial struct {
	RichGlobal
	CustomUser      int          `json:"customUser" gorm:"comment:所属客户"`
	Isp             int          `json:"isp" gorm:"type:int(1);default:1;comment:运营商"`
	Up              string       `json:"up" gorm:"default:0;comment:上行带宽"`
	Down            string       `json:"down" gorm:"default:0;comment:下行带宽"`
	Charging        int          `json:"charging" gorm:"type:int(1);default:0;comment:计费方式"`
	Price           float64      `json:"price" gorm:"comment:业务线单价"`
	ManagerLineCost float64      `json:"manager_line_cost" gorm:"comment:管理线价格"`
	BandwidthType   int          `json:"bandwidth_type" gorm:"default:0;comment:宽带类型"`
	TransProvince   bool         `json:"transProd" gorm:"default:false;comment:是否跨省"`
	MoreDialing     bool         `json:"moreDialing" gorm:"default:false;comment:是否支持多拨"`
	IsManager       int64        `json:"isManager" gorm:"type:tinyint(1);comment:是否管理线"`
	Account         string       `json:"account" gorm:"type:varchar(25);comment:账号"`
	Pass            string       `json:"pass" gorm:"type:varchar(20);comment:密码"`
	DialName        string       `json:"dialName" gorm:"type:varchar(20);comment:线路名称"`
	Status          int          `json:"status" gorm:"type:int(1);default:0;comment:拨号状态,1:已拨通 0:待使用 -1:拨号异常"`
	Source          int          `json:"source" gorm:"type:int(1);default:0;comment:拨号状态,0:录入 1:自动上报"`
	IdcId           int          `json:"idcId" gorm:"index;comment:关联的IDC"`
	HostId          int          `json:"hostId" gorm:"index;comment:关联主机ID"`
	DeviceId        int          `json:"deviceId" gorm:"index;comment:关联网卡ID"`
	RunTime         sql.NullTime `json:"runTime"`
}

func (Dial) TableName() string {
	return "rs_dial"
}
