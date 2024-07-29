package models

type RsCloud struct {
	RichGlobal
	Name string `json:"name" gorm:"index;type:varchar(50);comment:业务云名称"`
}

func (RsCloud) TableName() string {
	return "rs_cloud"
}
