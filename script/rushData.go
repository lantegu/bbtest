package datascript

import (
	storage "bbtest/storage"
)

// 对于一个编辑过滤器的使用者来说，他是否应该考虑RushDataWrapper是什么，有哪些参数

func FilterFeMale(rushDataWrapper storage.RushDataWrapper) (logger storage.Logger) {
	logger = storage.NewLocalLog(storage.Success, rushDataWrapper.Id, rushDataWrapper.LogInfo[0], nil)
	// for _, str := range row {
	// 	if str == "F" {
	// 		return logger, nil
	// 	}
	// }
	return logger
}