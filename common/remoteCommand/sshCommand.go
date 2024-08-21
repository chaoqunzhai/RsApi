package remoteCommand

import (
	"errors"
	"fmt"
	models2 "go-admin/cmd/migrate/migration/models"
	"go-admin/common/internal"
	"go-admin/config"
	"gorm.io/gorm"
	"os/exec"
	"regexp"
	"time"
)

// 借用ssh -t -o  -i 命令行执行shell

type Command struct {
	JobId      string        `json:"jobId"`
	Timeout    time.Duration `json:"timeout"`
	IdRsa      string        `json:"idRsa"`      //私钥路径
	RemotePort string        `json:"remotePort"` //frpc 建联端口号
	Orm        *gorm.DB      `json:"-"`
	HostId     int           `json:"hostId"`
	CreateBy   int           `json:"createBy"`
}

func (c Command) buildShell(shell string) (runShell string, err error) {
	matched, err := regexp.MatchString("rm", shell)

	if err != nil {
		return "", err
	}

	if matched {
		return "", errors.New("字符串中包含'rm'关键字")
	}

	//ssh -o Port=10302 -i  /root/.ssh/id_rsa  root@frp.xarscloud.com "hostname"
	remoteShell := fmt.Sprintf("%v %v \"%v\"", config.ExtConfig.Frps.IdRsa, config.ExtConfig.Frps.Address, shell)

	runShell = fmt.Sprintf("ssh -o StrictHostKeyChecking=no -o Port=%v -i %v",
		c.RemotePort, remoteShell)

	return runShell, nil
}

//执行命令

func (c Command) runShell(shell string) (output string, status int) {
	if c.RemotePort == "" {

		return fmt.Sprintf("failed host RemotePort is null"), -1
	}
	var err error
	var isShell string
	isShell, err = c.buildShell(shell)
	if err != nil {
		return fmt.Sprintf("buildShell %v", err), -1
	}

	cmd := exec.Command("bash", "-c", isShell)
	out, err := internal.CombinedOutputTimeout(cmd, 80*time.Second)
	if err != nil {
		return fmt.Sprintf("failed to run command %v: %v - %s", isShell, err, string(out)), -1
	}

	return string(out), 1
}

//业务切换

func (c Command) BusinessSwitching(bu string) (JobId int) {

	defaultBuPath := "/etc/business"

	shell := fmt.Sprintf("echo \"%v\" > %v", bu, defaultBuPath)

	output, status := c.runShell(shell)

	return c.SaveLog(status, output, "BusinessSwitching", shell)
}

//主机名修改

func (c Command) UpdateHostName(hostname string) (JobId int) {

	shell := fmt.Sprintf("hostnamectl set-hostname %v --static", hostname)

	output, status := c.runShell(shell)

	return c.SaveLog(status, output, "UpdateHostName", shell)
}

//执行命令

func (c Command) ExecuteCommand(shell string) (JobId int) {

	output, status := c.runShell(shell)

	return c.SaveLog(status, output, "ExecuteCommand", shell)
}

//重启机器

func (c Command) RebootHost() (JobId int) {

	runShell := "reboot"
	output, status := c.runShell(runShell)

	return c.SaveLog(status, output, "RebootHost", runShell)
}

func (c Command) SaveLog(status int, output, module, shell string) (JobId int) {
	var data models2.HostExecLog
	data.CreateBy = c.CreateBy
	data.HostId = c.HostId
	data.Exec = shell
	data.Module = module
	data.OutPut = output
	data.Status = status
	data.JobId = c.JobId
	c.Orm.Create(&data)

	return data.Id

}
