package storage

import (
	"time"
)

// RushDataWrapper wrap necessary Info about data
// include rowData (read by file), readTime and Id(read order)
type RushDataWrapper struct {
	Id        int
	readTime time.Time
	LogInfo []string
}

func NewRushDataWrapper(id int, logInfo []string, readTime time.Time) *RushDataWrapper {
	return &RushDataWrapper{Id:id, LogInfo: logInfo, readTime: readTime}
}

