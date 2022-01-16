package scripttool

import (
	"errors"
	"math/rand"
	"time"
)

func GetRowdataList(filepath string , readData func (filepath string) ([][]string, error),
dataReadPattern int, proportion int) (res [][]string, err error){
	batchData, err := readData(filepath)
	if err != nil {
		return nil, err
	}
	batchData, err = PatternFilter(batchData, dataReadPattern, proportion)
	if err != nil {
		return nil, err
	}
	return batchData, nil
}

func PatternFilter(batchData [][]string, dataReadPattern int, proportion int) ([][]string, error) {
	if  proportion < 0 || proportion > 100 {
		return nil, errors.New("proportion should between in 0-100")
	} 
	if dataReadPattern == Percent {
		return percentSelected(batchData, proportion), nil
	}
	return randomSelected(batchData, proportion), nil
}

func randomSelected(batchData [][]string, proportion int) (res [][]string) {
	res = make([][]string, 0)
	for _, data := range batchData {
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(100) + 1 <= proportion {
			res = append(res, data)
		}
	}
	return res
}

func percentSelected(batchData [][]string, proportion int) (res [][]string) {
	res = make([][]string, int(len(batchData) * proportion/100))
	copy(res, batchData[:int(len(batchData) * proportion/100)])
	return res
}