package scripttool

import (
	storage "bbtest/storage"
	"context"
	"time"
)

func TimeoutFunc(timeOut time.Duration, recordListStream chan <- storage.Logger,
	rowData storage.RushDataWrapper,
	rushData func(rushDataWrapper storage.RushDataWrapper) (storage.Logger)) {
	ctx, cancel := context.WithCancel(context.Background())
	var logger storage.Logger
	defer cancel()
	go func(ctx context.Context){
		logger = rushData(rowData)
		cancel()
	}(ctx)
	select {
	case <- ctx.Done():
		recordListStream <- logger
		return
	case <- time.After(timeOut):
		logger = storage.NewTimeOutLogger(rowData.Id, rowData.Info[0])
		recordListStream <- logger
		return
	}
}