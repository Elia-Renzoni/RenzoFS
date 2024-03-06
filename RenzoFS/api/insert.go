/**
*	@author Elia Renzoni
*	@date 02/02/2024
*	@brief Insert json query to files API
**/

package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
		writeInRemoteCSVFile(w)
	}
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

func writeInRemoteCSVFile(w http.ResponseWriter) {
	file, err := os.OpenFile("elia.csv", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		jsonErr, _ := errMessage.MarshalErrMessage(4)
		writeNegativeJSONResponse(w, jsonErr)
		log.Fatal(err)
	}
	defer file.Close()
	defer func() {
		for _, value := range payload.QueryContent {
			fmt.Printf("Value : %v", value)
		}
	}()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	if errOnWrite := writer.Write(payload.QueryContent); errOnWrite != nil {
		panic(err)
	}
	writeSuccessJSONResponse(w)
}
