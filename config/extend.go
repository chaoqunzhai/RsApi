package config

var ExtConfig Extend

// Extend 扩展配置
//
//	extend:
//	  demo:
//	    name: demo-name
//
// 使用方法： config.ExtConfig......即可！！
type Extend struct {
	AMap       AMap       // 这里配置对应配置文件的结构即可
	Prometheus prometheus `json:"prometheus"`
	Frps       Frps       `json:"frps"`
	Automation Automation `json:"automation"`
}
type Automation struct {
	Hostname string `json:"hostname"`
}

type Frps struct {
	Address string `json:"address"`
	IdRsa   string `json:"id_rsa"`
}
type prometheus struct {
	Endpoint string `json:"endpoint"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type AMap struct {
	Key string
}
