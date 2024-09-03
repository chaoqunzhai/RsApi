package models

type Custom struct {
	RichGlobal
	Name        string `json:"name" gorm:"type:varchar(120);comment:客户名称"`
	Type        int    `json:"type" gorm:"comment:客户类型,customer_type"`
	Cooperation int    `json:"cooperation" gorm:"comment:合作状态,work_status"`
	Region      string `json:"region" gorm:"type:varchar(200);comment:所在地区"`
	Address     string `json:"address"  gorm:"type:varchar(200);comment:地址"`
}

func (Custom) TableName() string {
	return "rs_custom"
}
