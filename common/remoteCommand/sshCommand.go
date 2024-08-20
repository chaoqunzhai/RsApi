package remoteCommand

import (
	"errors"
	"fmt"
	"go-admin/common/internal"
	"go-admin/config"
	"os/exec"
	"regexp"
	"time"
)

// 借用ssh -t -o  -i 命令行执行shell

type Command struct {
	Timeout    time.Duration `json:"timeout"`
	IdRsa      string        `json:"id_rsa"`      //私钥路径
	RemotePort int           `json:"remote_port"` //frpc 建联端口号
	Domain     string        `json:"domain"`      // frps 配置的域名
}

func (c Command) buildShell(sell string) (shell string, err error) {
	matched, err := regexp.MatchString("rm", sell)

	if err != nil {
		return "", err
	}

	if matched {
		return "", errors.New("字符串中包含'rm'关键字")
	}

	//ssh -o Port=10302 -i  /root/.ssh/id_rsa  root@frp.xarscloud.com "hostname"

	runShell := fmt.Sprintf("ssh -o Port=%v -i %v  %v \"%v\"",
		c.RemotePort, config.ExtConfig.Frps.IdRsa, config.ExtConfig.Frps.Address, sell)

	if c.Timeout == 0 {
		c.Timeout = 60 * time.Second
	}
	return runShell, nil
}

//执行命令

func (c Command) runShell(shell string) (output string, err error) {

	var isShell string
	isShell, err = c.buildShell(shell)
	if err != nil {
		return "", err
	}

	cmd := exec.Command(isShell)
	out, err := internal.CombinedOutputTimeout(cmd, c.Timeout)
	if err != nil {
		return fmt.Sprintf("failed to run command %v: %w - %s", isShell, err, string(out)), nil
	}

	return string(out), nil
}

//业务切换

func (c Command) SwitchBuShell(bu string) (output string, err error) {

	defaultBuPath := "/etc/business"

	shell := fmt.Sprintf("echo \"%v\" > %v", bu, defaultBuPath)

	return c.runShell(shell)
}

//主机名修改

func (c Command) UpdateHostName(hostname string) (output string, err error) {

	if hostname == "" {

		return "", errors.New("主机名不能为空")
	}

	if len(hostname) < 10 {

		return "", errors.New("主机名不得少于10个字符")
	}
	shell := fmt.Sprintf("hostnamectl set-hostname %v --static", hostname)

	return c.runShell(shell)
}

//重启服务

func (c Command) RestartServer(shell string) (output string, err error) {

	return c.runShell(shell)
}

//重启机器

func (c Command) RebootHost() (output string, err error) {

	return c.runShell("reboot")
}
