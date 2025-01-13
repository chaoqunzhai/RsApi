package models

type DataBurningHost struct {
	Model
	ModelTime
	Online int64 `json:"online" gorm:"online"` //在线数量
	Offline int64 `json:"offline" gorm:"offline"` //离线数量
	Wait int64 `json:"wait" gorm:"wait"` //待上架
	Todo int64 `json:"todo" gorm:"todo"` //已知晓
	TotalBandwidth  float64 `json:"total_bandwidth" gorm:"total_bandwidth"` //总带宽
	DiffHost string `json:"diff_host"  gorm:"diff_host"`  //今日计算 跟昨天相比有非在线状态的主机
}
func (DataBurningHost) TableName() string {
	return "rs_data_burning_host"
}