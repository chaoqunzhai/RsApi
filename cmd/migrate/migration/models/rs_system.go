package models

type Business struct {
	RichGlobal
	Name      string `json:"name" gorm:"index;type:varchar(50);comment:业务云名称"`
	Algorithm string `json:"algorithm" gorm:"type:varchar(120);comment:算法备注"`
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
