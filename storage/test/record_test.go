package storage

import (
	"testing"
	storage "bbtest/storage"
)
func TestCsvRecorder(t *testing.T) {
	csvRecoder, err := storage.NewCsvRecorder("./test.csv")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	testData := [][]string{{"hello","world","!"},
	{"you", "are", "good"},	
}
	for _, row := range testData {
		err = csvRecoder.RecordData(row)
		if err != nil {
			t.Log(err)
			t.Fail()
			return
		}
	}
	csvRecoder.Flush()
	csvRecoder.Close()
}