package models

// 联系人
type CustomUser struct {
	RichGlobal
	UserName string `json:"username" gorm:"type:varchar(50);comment:姓名"`
	CustomId int    `json:"customId" gorm:"index;comment:所属客户"`
	BuId     int    `json:"buId" gorm:"index;comment:所属商务人员"`
	Phone    string `json:"phone" gorm:"type:varchar(20);comment:联系号码"`
	Email    string `json:"email" gorm:"type:varchar(50);comment:联系邮箱"`
	Region   string `json:"region" gorm:"type:varchar(100);comment:省份城市多ID"`
	Dept     string `json:"dept" gorm:"type:varchar(30);comment:部门"`
	Duties   string `json:"duties" gorm:"type:varchar(30);comment:职务"`
	Address  string `json:"address" gorm:"type:varchar(255);comment:详细地址"`
}

func (CustomUser) TableName() string {
	return "rs_custom_user"
}
