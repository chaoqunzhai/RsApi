/*
*
@Author: chaoqun
* @Date: 2024/7/25 22:49
*/
package dto

type IdcMetrics struct {
	Remark   string
	Province string
	City     string
}
type RegisterMetrics struct {
	Belong         int               `json:"belong"`
	Version        string            `json:"version"`
	CPU            int               `json:"CPU"`
	Memory         uint64            `json:"memory"`
	Disk           string            `json:"disk"`
	Sn             string            `json:"sn"`
	Hostname       string            `json:"hostname"`
	Ip             string            `json:"ip"`
	Mac            string            `json:"mac"`
	Mask           string            `json:"mask"`
	Gateway        string            `json:"gateway"`
	NetType        int               `json:"netType"`
	PublicIp       string            `json:"publicIp"`
	Business       string            `json:"business"`
	Kernel         string            `json:"kernel"`
	RemotePort     string            `json:"remotePort"`
	BusinessSn     map[string]string `json:"business_sn"`
	Remark         string            `json:"remark"`
	Province       string            `json:"province"`
	City           string            `json:"city"`
	Isp            string            `json:"isp"`
	NetDevice      string            `json:"netDevice"`
	Balance        float64           `json:"balance"`
	BandwidthCnf   BandWithCnf       `json:"bandwidthCnf"`
	TransmitNumber float64           `json:"transmitNumber"`
	ReceiveNumber  float64           `json:"receiveNumber"`
	MemoryMap      map[string]uint64 `json:"memoryMap"`
	Dial           []*RegisterDial   `json:"dial"` //拨号列表
	ExtendMap      []SoftwareRow     `json:"extendMap"`

	//七牛的请求
	Node         string            `json:"node"`
	NatType      string            `json:"natType"`
	DialType     string            `json:"dialType"`
	DialSource   string            `json:"dialSource"`
	IsMultiLine  interface{}       `json:"isMultiLine"`
	Usbw         int               `json:"usbw"`
	BwNum        int               `json:"bwNum"`
	Bandwidth    int               `json:"bandwidth"`
	QnDial       []*RegisterQnDial `json:"qnDial"`
	QnAssetDisk  []*QnAssetDisk    `json:"qnAssetDisk"`
	QnInterfaces []*QnInterfaces   `json:"qnInterfaces"`
}
type BandWithCnf struct {
	Line  float64 `json:"line"`
	Width float64 `json:"width"`
}

type QnInterfaces struct {
	NetDevName string `json:"netDevName"`
	Ip         string `json:"ip"`
	Mac        string `json:"mac"`
	Speed      string `json:"speed"`
}
type SoftwareRow struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Desc  string `json:"desc"`
}
type NiuLinkMetrics struct {
	Data     []RegisterMetrics `json:"data"`
	Business string            `json:"business"`
}

type RegisterQnDial struct {
	NetDevName     string                     `json:"netDevName"`
	Ip             string                     `json:"ip"`
	Speed          string                     `json:"speed"`
	Type           string                     `json:"type"`
	DialStatusInfo []RegisterQnDialStatusInfo `json:"dialStatusInfo"`
}
type RegisterQnDialStatusInfo struct {
	Account           string `json:"account"`
	Password          string `json:"password"`
	VlanId            int    `json:"vlanId"`
	AdslNum           int    `json:"adslNum"`
	Ip                string `json:"ip"`
	DialStatus        string `json:"dialStatus"`
	ConnectStatus     string `json:"connectStatus"`
	Ipv6              string `json:"ipv6"`
	Ipv6ConnectStatus string `json:"ipv6ConnectStatus"`
	Bras              string `json:"bras"`
	PppInterface      string `json:"pppInterface"`
	Mac               string `json:"mac"`
	DialLog           string `json:"dialLog"`
}

type QnAssetDisk struct {
	DiskName        string  `json:"diskName"`
	Sn              string  `json:"sn"`
	IsSystem        bool    `json:"isSystem"`
	Type            string  `json:"type"`
	Size            int64   `json:"size"`
	Usage           float64 `json:"usage"`
	WIops           int     `json:"wIops"`
	RIops           int     `json:"rIops"`
	DiskMeasureInfo struct {
		MeasureCost int    `json:"measureCost"`
		StartTime   string `json:"startTime"`
		State       string `json:"state"`
	} `json:"diskMeasureInfo"`
	OccupantStatus bool `json:"occupantStatus"`
}

type DiskFields struct {
	Dev  string `json:"dev"`
	FS   string `json:"fs"`
	U    string `json:"u"` //用户使用
	T    string `json:"t"` //总量
	M    string `json:"m"` //挂载点
	UP   string `json:"up"`
	Type string `json:"type"`
}

type DianXinMetrics struct {
	Business string               `json:"business"`
	Data     []*DianXinMetricsRow `json:"data"`
}
type DianXinMetricsRow struct {
	Sn            string          `json:"sn"`
	Remark        string          `json:"remark"`
	SystemRoot    DxSystemRoot    `json:"system_root"`
	MountDisk     []*DxDisk       `json:"mount_disk"`
	Memory        *DxMemory       `json:"memory"`
	IspName       string          `json:"ispName"`
	Province      string          `json:"province"`
	City          string          `json:"city"`
	IsProvince    string          `json:"isProvince"`
	DxInterfaces  []*DxInterfaces `json:"interfaces"`
	Price         *DxPrice        `json:"price"`
	DialList      []*DxDialList   `json:"dial_list"`
	AllLine       int             `json:"allLine"`
	DeviceInfo    DeviceInfo      `json:"deviceInfo"`
	Balance       float64         `json:"balance"`
	LineBandwidth int             `json:"lineBandwidth"`
}
type DeviceInfo struct {
	MemTotal uint64 `json:"memTotal"`
	CpuCore  int    `json:"cpuCore"`
}
type DxInterfaces struct {
	IfName string `json:"ifName"`
	NicMac string `json:"nicMac"`
}
type DxDialList struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	VlanId   string `json:"vlanId"`
	VlanIp   string `json:"vlanIp"`
	Name     string `json:"name"`
	IfName   string `json:"ifName"`
	Status   string `json:"status"`
}
type DxMemory struct {
	CpuUsedRate float64 `json:"cpu_used_rate"`
	MemUsedRate float64 `json:"mem_used_rate"`
}
type DxDisk struct {
	Name  string  `json:"name"`
	Sn    string  `json:"sn"`
	Type  string  `json:"type"`
	Total float64 `json:"total"`
	Used  float64 `json:"used"`
	Ug    string  `json:"ug"`
}
type DxSystemRoot struct {
	Total float64 `json:"total"`
	Used  float64 `json:"used"`
	Ug    string  `json:"ug"`
}
type DxPrice struct {
	Price           int     `json:"price"`
	Amount          float64 `json:"amount"`
	TotalBandwidth  float64 `json:"totalBandwidth"`
	ResendBandwidth float64 `json:"resendBandwidth"`
	Bandwidth       float64 `json:"bandwidth"`
	WorkDate        string  `json:"workDate"`
	SettleStatus    string  `json:"settleStatus"`
}
