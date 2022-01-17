package scripttool

import (
	"fmt"
	script "bbtest/script"
	scripttool "bbtest/scripttool"
	"bbtest/storage"
	"testing"
)

func TestChanStringList(t *testing.T) {
	stringListStream := make(chan []string)

	go func () {
		defer close(stringListStream)
		stringListCase1 := []string{"hello","world"}
		stringListCase2 := []string{"hello","world","!"}
		stringListCase3 := make([]string, 0)
		stringListStream <- stringListCase1
		stringListStream <- stringListCase2
		stringListStream <- stringListCase3
	}()
	for stringList := range stringListStream {
		t.Log(stringList)
	}
}


func TestRunSh(t *testing.T) {
	funcRun := scripttool.NewFuncRun("./logger")
	recorder, err := storage.NewCsvRecorder("./test.csv")
	if err != nil {
		fmt.Print(err)
	}
	err = funcRun.RunSh("./heart.csv", recorder, scripttool.Percent, 100,script.ReadCsvFile, script.FilterOlder)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("rush data done")
	}
}
