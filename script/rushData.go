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
			// 某一行非年龄数据，输出失败日志
			return storage.NewLocalLog(storage.Failed, rushDataWrapper.Id, rushDataWrapper.Info[0], err)
		}
		if age <= 60 {
			// 小于60输出成功
			return  storage.NewLocalLog(storage.Success, rushDataWrapper.Id, rushDataWrapper.Info[0], nil)
		}
	}
	// 大于60岁输出失败。
	return storage.NewLocalLog(storage.Failed, rushDataWrapper.Id, rushDataWrapper.Info[0], nil)
}