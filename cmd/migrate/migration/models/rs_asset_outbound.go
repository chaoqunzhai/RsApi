package models

type OutboundOrder struct {
	RichGlobal
	Code           string `json:"code"  gorm:"type:varchar(50);comment:出库编码" `
	CustomId       int    `json:"customId" gorm:"type:bigint;comment:所属客户ID"`
	PhoneNumber    string `json:"phoneNumber" gorm:"type:varchar(20)"`
	Region         string `json:"region" gorm:"type:varchar(100);comment:省份城市多ID"`
	Ems            string `json:"ems" gorm:"type:varchar(20)"`
	TrackingNumber string `json:"trackingNumber" gorm:"type:varchar(30);comment:物流单号"`
	Address        string `json:"address" gorm:"type:varchar(255);comment:联系地址"`
	User           string `json:"user" gorm:"type:varchar(50);comment:联系人"`
	IdcId          int    `json:"idcId" gorm:"type:bigint;comment:idcId"`
}

func (OutboundOrder) TableName() string {
	return "asset_outbound_order"
}
