package apis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"github.com/google/uuid"
	"go-admin/app/cmdb/models"
	models2 "go-admin/cmd/migrate/migration/models"

	"go-admin/common/dto"
	"go-admin/common/remoteCommand"
)

type ExecBuSwitch struct {
	HostIds []int `json:"hostIds"` //多个主机
	BuIds   []int `json:"buIds"`   //切换的业务
}

type ExecUpHostNameReq struct {
	HostId     int    `json:"hostId"`     //多个主机
	HostName   string `json:"hostName"`   //主机名
	Automation int    `json:"automation"` //1:自动化 2:自定义
}
type ExecUpHostDescReq struct {
	HostId     int    `json:"hostId"`     //多个主机
	Desc   string `json:"desc"`   //主机名
}
type SuspendUpHostDescReq struct {
	HostId     int    `json:"hostId"`     //多个主机
	Desc   string `json:"desc"`   //主机名
	Tag bool `json:"tag"`
}
type ExecCommandReq struct {
	HostIds []int  `json:"hostIds"` //多个主机
	Shell   string `json:"shell"`   //执行的命令
}
type ExecRebootReq struct {
	HostIds []int `json:"hostIds"` //多个主机
}
type JobGetReq struct {
	JobId string `uri:"jobId"`
}

type GetHostReq struct {
	Id             int    `uri:"id" search:"-"`
	Status         int    `form:"status" search:"type:exact;column:status;table:rs_host_exec_log" comment:"状态"`
	Module         string `form:"module"  search:"type:exact;column:module;table:rs_host_exec_log" comment:"模块"`
	dto.Pagination `search:"-"`
}

func (m *GetHostReq) GetNeedSearch() interface{} {
	return *m
}


func (e RsHost) UpHostDesc(c *gin.Context) {
	req := ExecUpHostDescReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}


	var hostModel models.RsHost
	e.Orm.Model(&models.RsHost{}).Where("id = ?", req.HostId).Limit(1).Find(&hostModel)

	if hostModel.Id == 0 {
		e.Error(500, nil, "主机不存在")
		return
	}

	descModel :=&models.HostDesc{
	}

	if hostModel.Desc != ""{
		json.Unmarshal([]byte(hostModel.Desc),&descModel)

	}
	//只赋值设置的字段
	descModel.Desc = req.Desc
	descValue,_:=json.Marshal(&descModel)
	e.Orm.Model(&models.RsHost{}).Where("id = ?", req.HostId).Updates(map[string]interface{}{
		"desc":string(descValue),
	})
	e.OK("successful", "")
	return
}

func (e RsHost) ExecUpHostName(c *gin.Context) {
	req := ExecUpHostNameReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var Info string

	//自定义内容
	if req.Automation == 2 {
		if req.HostName == "" {
			err = errors.New("主机名不能为空")
			e.Error(500, err, err.Error())
			return
		}

		if len(req.HostName) < 10 {

			err = errors.New("主机名不得少于10个字符")
			e.Error(500, err, err.Error())
			return
		}
		Info = fmt.Sprintf("远程更新主机名:%v", req.HostName)
	} else {
		Info = "自动化生成主机名"
	}

	var hostModel models.RsHost
	e.Orm.Model(&models.RsHost{}).Where("id = ?", req.HostId).Limit(1).Find(&hostModel)

	if hostModel.Id == 0 {
		e.Error(500, nil, "主机不存在")
		return
	}
	JobId := uuid.New().String()
	command := remoteCommand.Command{
		Orm:        e.Orm,
		CreateBy:   user.GetUserId(c),
		HostId:     req.HostId,
		RemotePort: hostModel.RemotePort,
		JobId:      JobId,
	}
	go func() {
		command.UpdateHostName(req.Automation, req.HostName)
	}()

	e.Orm.Create(&models2.OperationLog{
		CreateUser: user.GetUserName(c),
		Action:     "POST",
		Module:     "rs_host",
		ObjectId:   hostModel.Id,
		TargetId:   hostModel.Id,
		Info:       Info,
	})
	e.OK(JobId, "")
	return
}

func (e RsHost) ExecCommand(c *gin.Context) {
	req := ExecCommandReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//批量执行命令
	var hostList []models.RsHost
	e.Orm.Model(&models.RsHost{}).Where("id in ?", req.HostIds).Find(&hostList)

	JobId := uuid.New().String()
	if len(hostList) == 0 {
		e.Error(500, nil, "主机为空")
		return
	}
	for _, host := range hostList {

		command := remoteCommand.Command{
			Orm:        e.Orm,
			CreateBy:   user.GetUserId(c),
			HostId:     host.Id,
			RemotePort: host.RemotePort,
			JobId:      JobId,
		}
		go func() {
			command.ExecuteCommand(req.Shell)
		}()

		e.Orm.Create(&models2.OperationLog{
			CreateUser: user.GetUserName(c),
			Action:     "POST",
			Module:     "rs_host",
			ObjectId:   host.Id,
			TargetId:   host.Id,
			Info:       fmt.Sprintf("远程执行命令:%v", req.Shell),
		})

	}

	e.OK(JobId, "")
	return
}



func (e RsHost) SuspendBilling(c *gin.Context) {
	req := SuspendUpHostDescReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}


	var hostModel models.RsHost
	e.Orm.Model(&models.RsHost{}).Where("id = ?", req.HostId).Limit(1).Find(&hostModel)

	if hostModel.Id == 0 {
		e.Error(500, nil, "主机不存在")
		return
	}

	descModel :=&models.HostDesc{
	}

	if hostModel.Desc != ""{
		marErr:=json.Unmarshal([]byte(hostModel.Desc),&descModel)
		if marErr!=nil{
			descModel.Desc = hostModel.Desc
		}
	}
	descModel.SuspendBilling = req.Desc
	descValue,_:=json.Marshal(&descModel)
	e.Orm.Model(&models.RsHost{}).Where("id = ?", req.HostId).Updates(map[string]interface{}{
		"desc":string(descValue),
		"suspend_billing":req.Tag,
	})
	e.Orm.Create(&models.RsHostSuspendLog{
		HostId: int64(req.HostId),
		CreateBy: user.GetUserId(c),
		Desc:  req.Desc,
		Enable: req.Tag,
	})
	e.OK("successful", "")
	return
}
func (e RsHost) ExecReboot(c *gin.Context) {
	req := ExecRebootReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//批量执行命令
	var hostList []models.RsHost
	e.Orm.Model(&models.RsHost{}).Where("id in ?", req.HostIds).Find(&hostList)
	if len(hostList) == 0 {
		e.Error(500, nil, "主机为空")
		return
	}
	JobId := uuid.New().String()
	for _, host := range hostList {

		command := remoteCommand.Command{
			Orm:        e.Orm,
			CreateBy:   user.GetUserId(c),
			HostId:     host.Id,
			RemotePort: host.RemotePort,
			JobId:      JobId,
		}
		go func() {
			command.RebootHost()
		}()
		e.Orm.Create(&models2.OperationLog{
			CreateUser: user.GetUserName(c),
			Action:     "POST",
			Module:     "rs_host",
			ObjectId:   host.Id,
			TargetId:   host.Id,
			Info:       "远程重启机器",
		})

	}
	e.OK(JobId, "")
	return
}

func (e RsHost) GetJobLog(c *gin.Context) {
	req := JobGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var data []models2.HostExecLog
	e.Orm.Model(&data).Where("job_id = ?", req.JobId).Find(&data)

	if len(data) == 0 {
		e.OK(map[string]string{
			"outPut": "执行中......\n\n",
		}, "")
		return
	}
	var outputAll string
	for _, row := range data {
		outputAll += row.OutPut + "\n"
	}
	firstRow := data[0]
	firstRow.OutPut = outputAll
	if firstRow.OutPut == "" {
		firstRow.OutPut = "执行中......\n\n"
	}
	e.OK(firstRow, "")
	return
}

func (e RsHost) GetHostLog(c *gin.Context) {
	req := GetHostReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var data []models2.HostExecLog
	e.Orm.Model(&data).Where("host_id = ?", req.Id).Scopes(
		dto.MakeCondition(req.GetNeedSearch()),
		dto.Paginate(req.GetPageSize(), req.GetPageIndex())).Order("id desc").Find(&data)

	e.PageOK(data, len(data), req.GetPageIndex(), req.GetPageSize(), "查询成功")
	return
}
