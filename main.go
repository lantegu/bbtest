package main

import (
	script "bbtest/script"
	scripttool "bbtest/scripttool"
	"bbtest/storage"
	"fmt"
)

func main (){
	funcRun := scripttool.NewFuncRun("./logger")
	recorder, err := storage.NewCsvRecorder("./test.csv")
	if err != nil {
		fmt.Print(err)
	}
	err = funcRun.RunSh("./scripttool/test/sub_heart.csv", recorder, script.ReadCsvFile, script.FilterOlder)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("rush data done")
	}
}