package scripttool

import (
	"context"
	storage "bbtest/storage"
	"sort"
	"sync"
)

// 多信号量模式，开启时执行多个rushData协程，每个完成后给logDataRoutine发送完成信号，收到所有完成信号后
//logData关闭

func rushDataRoutine(ctx context.Context, routineNum int, rushDataListStream <-chan storage.RushDataWrapper,
	rushData func(rushDataWrapper storage.RushDataWrapper) (storage.Logger)) chan storage.Logger  {
		recordListStream := make(chan storage.Logger)
		go func() {
			defer close(recordListStream)
			for rowData := range rushDataListStream {
				select {
				case <- ctx.Done():
					return 
				default:
				}
				//此处设置超时模块，
				logger := rushData(rowData)
				recordListStream <- logger
			}
		}()
		return recordListStream
}

// 添加一个功能，将所有的logger信号集合起来，然后用recorder进行一个存储,
// 多信号量模式
func logDataRoutine(ctx context.Context, loggerListStream <-chan storage.Logger, recorder storage.Recorder, 
	done chan<- interface{})  {
	defer func (){
		recorder.Close()
		done <- 1
	}()
	logDataList := make([]storage.Logger, 0)
	for logData := range loggerListStream {
		select {
		case <-ctx.Done():
			return
		default:
		}
		logDataList = append(logDataList, logData)
	}
	sort.Slice(logDataList,func(i,j int) bool {
		return logDataList[i].GetId() < logDataList[j].GetId()
	})
	for _, logger := range logDataList {
		recorder.RecordData(logger.LogData())
	}
	recorder.Flush()
	return
}

// 扇入函数

func fanIn(channels ... chan storage.Logger) <-chan storage.Logger {
	var wg sync.WaitGroup
	multiplexStream := make(chan storage.Logger)
	multiplex := func(c <-chan storage.Logger, idNum int) {
		defer wg.Done()
		for i := range c {
			multiplexStream <- i
		}
	}
	wg.Add(len(channels))
	for i, c := range channels {
		go multiplex(c, i)
	}

	go func() {
		wg.Wait()
		close(multiplexStream)
	}()
	return multiplexStream
}
