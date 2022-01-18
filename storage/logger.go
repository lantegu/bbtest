package storage

import (
	"errors"
	"fmt"
)

// 本地和云端都做记录即可,这里只是编译数据log信息
type Logger interface {
	GetId() int
	LogData() []string
}

// 目前就先接受 []string, 以后要有拓展需要用s
type LocalLog struct {
	status int
	primaryKey string
	id int
	err error
}

func (localLog LocalLog) GetId() int {
	return localLog.id
}

func (localLog LocalLog) LogData() []string {
	if localLog.status == Failed && localLog.err != nil {
		return []string{localLog.primaryKey, fmt.Sprintf("update error:%v", localLog.err)}
	}
	if localLog.status == Failed && localLog.err == nil {
		return []string{localLog.primaryKey, "update error and not throw error info"}
	}
	return []string{localLog.primaryKey, "update success"}
} 

func NewLocalLog(status int, id int, primaryKey string, err error) LocalLog {
	return LocalLog{status: status, primaryKey: primaryKey,id : id, err: err}
}

func NewTimeOutLogger(id int, primaryKey string) LocalLog {
	return LocalLog{status: Failed, primaryKey: primaryKey, id : id, err: errors.New("timeout")}
}