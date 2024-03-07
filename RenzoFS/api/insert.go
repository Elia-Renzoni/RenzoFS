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
	"path/filepath"
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
	for {
		if err := os.Chdir(filepath.Join("local_file_system", payload.User)); err != nil {
			if os.IsNotExist(err) {
				controller.createNewDir(payload.User)
				fmt.Printf("directory creata")
			}
		} else {
			break
		}
	}

	file, err := os.OpenFile(payload.FileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		jsonErr, _ := errMessage.MarshalErrMessage(4)
		writeNegativeJSONResponse(w, jsonErr)
		log.Fatal(err)

		// if the file dont exist
		if os.IsNotExist(err) {
			jsonErr, _ := errMessage.MarshalErrMessage(4)
			writeNegativeJSONResponse(w, jsonErr)
		}
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	if errOnWrite := writer.Write(payload.QueryContent); errOnWrite != nil {
		panic(err)
	}
	writeSuccessJSONResponse(w)
}
