/**
*	@author Elia Renzoni
*	@date 02/02/2024
*	@brief Insert json query to files API
**/

package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type InsertPayLoad struct {
	User         string              `json:"user"`
	FileName     string              `json:"file_name"`
	QueryContent map[string][]string `json:"query_content"`
}

var (
	payload        InsertPayLoad      = InsertPayLoad{}
	errMessage     ResponseMessages   = ResponseMessages{}
	controller     ResourceController = ResourceController{}
	fileOperations FileOp             = FileOp{}
)

func HandleInsertion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json, _ := errMessage.MarshallErrMessage()
		w.Write(json)
	} else {
		defer r.Body.Close()
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &payload)
		if jsonMessage, _ := parseJSONQuery(); jsonMessage != nil {
			writeNegativeJSONResponse(w, jsonMessage)
		}
		insertQueryValuesToCSV()
	}
}

func parseJSONQuery() ([]byte, error) {
	if result := controller.checkDir(payload.User); result {
		return errMessage.MarshallDirException()
	}
	if result := controller.checkFile(payload.FileName); result {
		return errMessage.MarshallFileException()
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
	json, _ := errMessage.MarshallPOSTSuccess()
	w.Write(json)
}

func insertQueryValuesToCSV() {
	// lock
	fileOperations.writeCSV(payload.FileName, payload.QueryContent)
	// unlock
}

/*
func printContent() {
	fmt.Printf("User : %s\n", payload.User)
	fmt.Printf("FileName: %s\n", payload.FileName)
	for key, value := range payload.QueryContent {
		fmt.Printf("Key : %s", key)
		switch eff := value.(type) {
		case string:
			fmt.Printf("Value : %s", eff)
		case int:
			fmt.Printf("Value : %d", eff)
		case float64:
			fmt.Printf("Value: %f", eff)
		}
	}
}*/
