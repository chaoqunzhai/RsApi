package models

type Business struct {
	RichGlobal
	Status    int    `json:"status" gorm:"type:int(1);default:0;comment:拨号状态,1:正常 非1:异常"`
	Name      string `json:"name" gorm:"index;type:varchar(50);comment:业务云名称"`
	EnName    string `json:"en_name" gorm:"index;type:varchar(30);comment:业务英文名字"`
	Algorithm string `json:"algorithm" gorm:"type:varchar(120);comment:算法备注"`
	Host      []Host `json:"-" gorm:"many2many:host_bind_business;foreignKey:id;joinForeignKey:host_id;references:id;joinReferences:business_id;"`
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
