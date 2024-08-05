package models

type Dial struct {
	RichGlobal
	Number   string `json:"number" gorm:"type:varchar(60);comment:账号"`
	User     string `json:"user" gorm:"type:varchar(30);comment:用户名"`
	Pass     string `json:"pass" gorm:"type:varchar(50);comment:密码"`
	Status   int    `json:"status" gorm:"type:int(1);default:0;comment:拨号状态,1:正常 非1:异常"`
	IdcId    int    `json:"idcId" gorm:"index;comment:关联的IDC"`
	HostId   int    `json:"hostId" gorm:"index;comment:关联主机ID"`
	DeviceId int    `json:"deviceId" gorm:"index;comment:关联网卡ID"`
}

func (Dial) TableName() string {
	return "rs_dial"
}
