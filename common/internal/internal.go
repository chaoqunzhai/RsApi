package internal

import "errors"

var (
	ErrTimeout = errors.New("命令执行超时")
)
