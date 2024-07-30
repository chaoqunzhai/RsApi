package models

//机房信息

type Idc struct {
	RichGlobal
	Number          int     `json:"number" gorm:"comment:机房编号"`
	Name            string  `json:"name" gorm:"type:varchar(100);default:'';comment:机房名称"`
	CustomUser      int     `json:"customUser" gorm:"comment:所属客户"`
	Region          string  `json:"region" gorm:"type:varchar(50);comment:所在地区"`
	Address         string  `json:"address" gorm:"type:varchar(255);default:'';comment:详细地址"`
	IpV6            bool    `json:"IpV6" gorm:"default:false;comment:是否IPV6"`
	TypeId          int     `json:"type_id" gorm:"type:int(1);default:0;comment:机房类型"`
	BusinessUser    int     `json:"businessUser" gorm:"comment:商务人员"`
	WechatName      string  `json:"wechatName" gorm:"type:varchar(100);comment:企业微信群名称"`
	WebHookUrl      string  `json:"webHookUrl" gorm:"type:varchar(100);comment:企业微信webhookUrl"`
	Status          int     `json:"status" gorm:"type:int(1);default:1;comment:机房状态"`
	Belong          int     `json:"belong" gorm:"type:int(1);default:0;comment:机房归属"`
	Isp             int     `json:"isp" gorm:"type:int(1);default:1;comment:运营商"`
	AllBandwidth    string  `json:"all_bandwidth" gorm:"type:varchar(35);comment:机房总带宽"`
	AllLine         int     `json:"all_line" gorm:"type:int(1);default:0;comment:机房总线路"`
	Up              string  `json:"up" gorm:"default:0;comment:上行带宽"`
	Down            string  `json:"down" gorm:"default:0;comment:下行带宽"`
	Price           float64 `json:"price" gorm:"comment:单价"`
	ManageLine      int     `json:"manage_line" gorm:"type:int(1);default:0;comment: 管理线路数"`
	ManagerLineCost float64 `json:"manager_line_cost" gorm:"comment:管理线价格"`
	BandwidthType   int     `json:"bandwidth_type" gorm:"default:0;comment:宽带类型"`
	Charging        int     `json:"charging" gorm:"type:int(1);default:0;comment:计费方式"`
	TransProvince   bool    `json:"transProd" gorm:"default:false;comment:是否跨省"`
	MoreDialing     bool    `json:"moreDialing" gorm:"default:false;comment:是否支持多拨"`
}

func (Idc) TableName() string {
	return "rs_idc"
}
