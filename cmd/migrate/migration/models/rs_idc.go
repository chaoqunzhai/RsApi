package models

//机房信息

type Idc struct {
	RichGlobal
	Number        int    `json:"number" gorm:"index;unique;comment:机房编号"`
	Name          string `json:"name" gorm:"type:varchar(120);default:'';comment:机房名称"`
	BuId          int    `json:"buId" gorm:"type:bigint;comment:商务人员"`
	CustomId      int    `json:"customId" gorm:"type:bigint;comment:所属客户ID"`
	Region        string `json:"region" gorm:"type:varchar(120);comment:所在地区"`
	Address       string `json:"address" gorm:"type:varchar(255);default:'';comment:详细地址"`
	IpV6          bool   `json:"IpV6" gorm:"default:true;comment:是否IPV6"`
	TypeId        int    `json:"type_id" gorm:"type:int(1);default:3;comment:机房类型"`
	WechatName    string `json:"wechatName" gorm:"type:varchar(120);comment:企业微信群名称"`
	WebHookUrl    string `json:"webHookUrl" gorm:"type:varchar(200);comment:企业微信webhookUrl"`
	Status        int    `json:"status" gorm:"type:int(1);default:1;comment:机房状态"`
	Belong        int    `json:"belong" gorm:"type:int(1);default:0;comment:机房归属"`
	TransProvince bool   `json:"transProd" gorm:"default:false;comment:是否跨省"`
	MoreDialing   bool   `json:"moreDialing" gorm:"default:false;comment:是否支持多拨"`
}

func (Idc) TableName() string {
	return "rs_idc"
}
