package models

//机房信息

type Idc struct {
	RichGlobal

	Name          string  `json:"name" gorm:"type:varchar(100);default:'';comment:机房名称"`
	Status        int     `json:"status" gorm:"type:int(1);default:1;comment:机房状态"`
	Belong        int     `json:"belong" gorm:"type:int(1);default:0;comment:机房归属"`
	TypeId        int     `json:"type_id" gorm:"type:int(1);default:0;comment:机房类型"`
	BusinessUser  int     `json:"businessUser" gorm:"comment:商务人员"`
	CustomUser    int     `json:"customUser" gorm:"comment:所属客户"`
	Region        string  `json:"region" gorm:"type:varchar(50);comment:省份城市多ID"`
	Charging      int     `json:"charging" gorm:"type:int(1);default:0;comment:计费方式"`
	Price         float64 `json:"price" gorm:"comment:单价"`
	WeChatNumber  string  `json:"weChatNumber" gorm:"type:varchar(10);default:'';comment:企业微信群号"`
	IpV6          bool    `json:"IpV6" gorm:"default:false;comment:是否IPV6"`
	TransProvince bool    `json:"transProd" gorm:"default:false;comment:是否跨省"`
	Address       string  `json:"address" gorm:"type:varchar(255);default:'';comment:详细地址"`
	WebHookUrl    string  `json:"webHookUrl" gorm:"type:varchar(100);comment:企业微信webhookUrl"`
}

func (Idc) TableName() string {
	return "rs_idc"
}

//机房宽带

type IdcBandwidth struct {
	RichGlobal
	IdcId           int     `json:"idcId" gorm:"index;comment:机房ID"`
	All             float64 `json:"all" gorm:"default:0;comment:总带宽"`
	AllLine         int     `json:"all_line" gorm:"type:int(1);default:0;comment:总线路"`
	ManageLine      int     `json:"manage_line" gorm:"type:int(1);default:0;comment: 管理线路数"`
	ManagerLineCost float64 `json:"manager_line_cost" gorm:"comment:管理线价格"`
	Up              string  `json:"up" gorm:"default:0;comment:上行带宽"`
	Down            string  `json:"down" gorm:"default:0;comment:下行带宽"`
	BandType        int     `json:"bandType" gorm:"default:0;comment:宽带类型"`
	MoreDialing     bool    `json:"moreDialing" gorm:"default:false;comment:是否推荐"`
}

func (IdcBandwidth) TableName() string {
	return "rs_idc_bandwidth"
}
