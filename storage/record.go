package storage

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Recorder interface {
	RecordData(logRow []string) error
	Flush()
	Close()
}

type CsvRecorder struct {
	saveFilePath string
	csvFile   *os.File
	csvWriter *csv.Writer
}

func NewCsvRecorder(saveFilePath string) (*CsvRecorder, error) {
	csvRecoder := &CsvRecorder{saveFilePath:saveFilePath}
	err := csvRecoder.createWriter()
	if err != nil {
		return nil, err
	}
	return csvRecoder, nil
}

func (this *CsvRecorder) createWriter() (err error) {
	this.csvFile, err = os.Create(this.saveFilePath)
	if err != nil {
		return nil
	}
	this.csvWriter = csv.NewWriter(this.csvFile)
	return err
}

func (this *CsvRecorder) RecordData(logRow []string) error {
	err := this.csvWriter.Write(logRow)
	return err
}

func (this *CsvRecorder) Flush() {
	if this.csvWriter == nil {
		fmt.Printf("csv writer:%v is not create plase call method create writer", this.saveFilePath)
	}
	this.csvWriter.Flush()
}

func (this *CsvRecorder) Close() {
	if this.csvFile == nil {
		fmt.Printf("file:%v is not create plase call method create writer", this.saveFilePath)
		return
	}
	this.csvFile.Close()
}