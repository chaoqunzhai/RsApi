package models

type OutboundOrder struct {
	RichGlobal
	Code           string `json:"code"  gorm:"type:varchar(50);comment:出库编码/编码" `
	CustomId       int    `json:"customId" gorm:"type:bigint;comment:所属客户ID"`
	PhoneNumber    string `json:"phoneNumber" gorm:"type:varchar(20)"`
	Ems            string `json:"ems" gorm:"type:varchar(20);comment:物流公司"`
	TrackingNumber string `json:"trackingNumber" gorm:"type:varchar(30);comment:物流单号"`
	Address        string `json:"address" gorm:"type:varchar(255);comment:联系地址"`
	Region         string `json:"region" gorm:"type:varchar(100);comment:省份城市多ID"`
	UserId         int    `json:"userId" gorm:"comment:联系人"`
	IdcId          int    `json:"idcId" gorm:"type:bigint;comment:idcId"`
	Status         int    `json:"status" gorm:"index;type:int(1);default:1;comment:出库状态"`
}

func (OutboundOrder) TableName() string {
	return "asset_outbound_order"
}
