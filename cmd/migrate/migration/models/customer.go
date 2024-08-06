package models

type Customer struct {
	RichGlobal
	Name       string `json:"name" gorm:"type:varchar(20);comment:客户名称"`
	Region     string `json:"region" gorm:"type:varchar(80);comment:省份城市多ID"`
	Address    string `json:"address" gorm:"type:varchar(80);comment:地址"`
	Level      int    `json:"level" gorm:"type:int(1);default:0;comment:客户等级"`
	TypeId     int    `json:"typeId" gorm:"type:int(1);default:0;comment:客户类型"`
	WorkStatus int    `json:"workStatus" gorm:"type:int(1);default:0;comment:合作状态"`
}

func (Customer) TableName() string {
	return "rs_customer"
}

type Contacts struct {
	RichGlobal
	UserName   string `json:"userName" gorm:"type:varchar(30);comment:用户名"`
	CustomerId int    `json:"customerId" gorm:"type:int(1);default:0;comment:客户ID"`
	BuUser     int    `json:"buUser" gorm:"type:int(1);default:0;comment:商务人员"`
	Phone      string `json:"phone" gorm:"type:varchar(20);comment:电话号码"`
	Landline   string `json:"landline" gorm:"type:varchar(20);comment:座机号"`
	Region     string `json:"Region" gorm:"type:varchar(80);comment:管理区域,也是城市ID"`
	Email      string `json:"email" gorm:"type:varchar(20);comment:电话号码"`
	Address    string `json:"address" gorm:"type:varchar(100);comment:地址"`
	Department string `json:"department" gorm:"type:varchar(30);comment:部门"`
	Duties     string `json:"duties" gorm:"type:varchar(30);comment:职务"`
}

func (Contacts) TableName() string {
	return "rs_contacts"
}
