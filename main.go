package main

import (
	script "bbtest/script"
	scripttool "bbtest/scripttool"
	"bbtest/storage"
	"fmt"
	"time"
)

func main (){
	funcRun := scripttool.NewFuncRun("./logger")
	recorder, err := storage.NewCsvRecorder("./test.csv")
	if err != nil {
		fmt.Print(err)
	}
	options := scripttool.NewOptions(scripttool.Percent, 100, time.Millisecond, 5)
	err = funcRun.RunSh("./scripttool/test/sub_heart.csv", recorder, *options, script.ReadCsvFile, script.FilterOlder)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("rush data done")
	}
}