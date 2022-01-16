package storage

import "fmt"

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

func (this LocalLog) GetId() int {
	return this.id
}

func (this LocalLog) LogData() []string {
	if this.status == Failed && this.err != nil {
		return []string{this.primaryKey, fmt.Sprintf("update error:%v", this.err)}
	}
	if this.status == Failed && this.err == nil {
		return []string{this.primaryKey, fmt.Sprintf("update error and not throw error info")}
	}
	return []string{this.primaryKey, fmt.Sprintf("update success")}
} 

func NewLocalLog(status int, id int, primaryKey string, err error) LocalLog {
	return LocalLog{status: status, primaryKey: primaryKey,id : id, err: err}
}