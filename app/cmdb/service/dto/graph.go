package dto

type GraphPageReq struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	BuId      string `json:"buId"`
	IspId     string `json:"ispId"`
	HostId    string `json:"hostId"`
	CustomId  string `json:"customId"`
	IdcId     string `json:"idcId"`
}
