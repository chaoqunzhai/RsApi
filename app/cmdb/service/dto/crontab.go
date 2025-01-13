package dto

type DataBurningHost struct {
	Online int64 `json:"online"` //在线数量
	Offline int64 `json:"offline"` //离线数量
	Wait int64 `json:"wait"` //待上架
	Todo int64 `json:"todo"` //已知晓
	TotalBandwidth  float64 `json:"total_bandwidth"` //总带宽
	DiffHost []int `json:"diff_host"`  //今日计算 跟昨天相比有非在线的主机
}