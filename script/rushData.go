package datascript

import (
	storage "bbtest/storage"
	"strconv"
)

// 对于一个编辑过滤器的使用者来说，他是否应该考虑RushDataWrapper是什么，有哪些参数

func FilterOlder(rushDataWrapper storage.RushDataWrapper) (logger storage.Logger) {
	
	rowData := rushDataWrapper.Info
	if len(rowData) >=2 {
		age, err := strconv.Atoi(rowData[0])
		if err != nil {
			return storage.NewLocalLog(storage.Failed, rushDataWrapper.Id, rushDataWrapper.Info[0], err)
		}
		if age <= 60 {
			return  storage.NewLocalLog(storage.Success, rushDataWrapper.Id, rushDataWrapper.Info[0], nil)
		}
	}
	return storage.NewLocalLog(storage.Failed, rushDataWrapper.Id, rushDataWrapper.Info[0], nil)
}