package models

type RsCloud struct {
	RichGlobal
	Name string `json:"name" gorm:"index;comment:业务云类型"`
}

func (RsCloud) TableName() string {
	return "rc_cloud"
}
