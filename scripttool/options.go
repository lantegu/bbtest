package scripttool

import "time"

type Options struct {
	dataReadPattern int
	proportion      int
	timeOut         time.Duration
	retryCount      int
}

func NewOptions(dataReadPattern int, proportion int, timeOut time.Duration, retryCount int) *Options {
	return &Options{dataReadPattern: dataReadPattern, proportion: proportion, timeOut: timeOut,
	retryCount: retryCount}
}