/**
*	@author Elia Renzoni
*	@date 02/02/2024
*	@brief Insert json query to files API
**/

package api

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type InsertPayLoad struct {
	QueryType    string   `json:"query_type"`
	User         string   `json:"user"`
	FileName     string   `json:"file_name"`
	QueryContent []string `json:"query_content"`
}

var (
	payload    InsertPayLoad      = InsertPayLoad{}
	errMessage ResponseMessages   = ResponseMessages{}
	controller ResourceController = ResourceController{}
)

func HandleInsertion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json, _ := errMessage.MarshalErrMessage(0)
		w.Write(json)
	} else {
		defer r.Body.Close()
		reqBody, _ := io.ReadAll(r.Body)
		json.Unmarshal(reqBody, &payload)
		if jsonMessage, _ := parseJSONQuery(); jsonMessage != nil {
			writeNegativeJSONResponse(w, jsonMessage)
		}
		writeInRemoteCSVFile()
		writeSuccessJSONResponse(w)
	}
}

func parseJSONQuery() ([]byte, error) {
	if result := controller.checkDir(payload.User); result {
		return errMessage.MarshalErrMessage(3)
	}
	if result := controller.checkFile(payload.FileName); result {
		return errMessage.MarshalErrMessage(4)
	}
	return nil, nil
}

func writeNegativeJSONResponse(w http.ResponseWriter, jsonMessage []byte) {
	w.WriteHeader(http.StatusNotAcceptable)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonMessage)
}

func writeSuccessJSONResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	json, _ := errMessage.MarshalSuccessMessage(5)
	w.Write(json)
}

func writeInRemoteCSVFile() {
	file, err := os.Open("renzofs_local_file_system/" + payload.User + "/" + payload.FileName)

	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	if errOnWrite := writer.Write(payload.QueryContent); errOnWrite != nil {
		panic(err)
	}

}
