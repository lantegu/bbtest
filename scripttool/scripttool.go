package scripttool

import (
	"context"
	storage "bbtest/storage"
	"time"
)

// 多种模式，直接运行脚本，运行golang程序

type BaseInfo struct {
	startTime *time.Time
	endTime *time.Time
	outLocation string
}

type cmdRun struct {
	base *BaseInfo
}

type funcRun struct {
	base *BaseInfo
}

func NewCmdRun(outLocation string) *cmdRun {
	time := time.Now()
	return &cmdRun{base : &BaseInfo{startTime: &time,endTime: nil,
		outLocation: outLocation}}
}

func NewFuncRun(outLocation string) *funcRun {
	time := time.Now()
	return &funcRun{base : &BaseInfo{startTime: &time,endTime: nil,
		outLocation: outLocation}}
}

// 传入读取单位数据的函数,rushData返回一个接口，或者一个log错误
func (this *funcRun) RunSh(filepath string, recorder storage.Recorder, dataReadPattern int, proportion int,
readData func (filepath string) ([][]string, error), 
rushData func(rushDataWrapper storage.RushDataWrapper) (storage.Logger)) error {
	rowDataList, err := GetRowdataList(filepath, readData, dataReadPattern, proportion)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	rowDataStream := make(chan storage.RushDataWrapper)
	logDataStreamList := make([]chan storage.Logger, RoutineCount)
	for i := range logDataStreamList {
		logDataStreamList[i] = make(chan storage.Logger)
	}
	done := make(chan interface{})
	for i := 0; i < RoutineCount; i++ {
		logDataStreamList[i] = rushDataRoutine(ctx, i, rowDataStream, rushData)
	}

	defer func(){
		cancel()
	}()

	logDataStream := fanIn(logDataStreamList...)
	go logDataRoutine(ctx, logDataStream, recorder, done)

	for i, rowData := range rowDataList {
		rowDataStream <- *storage.NewRushDataWrapper(i, rowData, time.Now())  
	}
	close(rowDataStream)
	<- done
	return nil
}