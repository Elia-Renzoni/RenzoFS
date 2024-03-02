/**
*	@author Elia Renzoni
*	@date 02/03/2024
*	@brief
*
*
**/

package api

import (
	"encoding/csv"
	"os"
)

type FileOp struct {
}

func (f *FileOp) writeCSV(fileName string, queryContent map[string][]string) {
	file, _ := os.Open(fileName)

	writer := csv.NewWriter(file)

	for key, value := range queryContent {
		// TODO
	}
}

func (f *FileOp) readCSV() map[string]interface{} {
	return nil
}
