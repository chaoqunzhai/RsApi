package models

type Business struct {
	RichGlobal
	Status     int    `json:"status" gorm:"type:int(1);default:0;comment:拨号状态,1:正常 非1:异常"`
	Name       string `json:"name" gorm:"index;type:varchar(50);comment:业务云名称"`
	EnName     string `json:"enName" gorm:"index;type:varchar(30);comment:业务英文名字"`
	Algorithm  string `json:"algorithm" gorm:"type:varchar(30);comment:算法标记"`
	OpeMonitor bool   `json:"ope_monitor" gorm:"default:true;comment:是否支持业务监控"`
}

func (Business) TableName() string {
	return "rs_business"
}

type Tag struct {
	RichGlobal
	Name string `json:"name" gorm:"index;type:varchar(50);comment:业务云名称"`
}

func (Tag) TableName() string {
	return "rs_tag"
}
