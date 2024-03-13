/**
*	@author Elia Renzoni
*	@date 02/02/2024
*	@brief Insert json query to files API
**/

package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type InsertPayLoad struct {
	QueryType    string   `json:"query_type"`
	User         string   `json:"user"`
	FileName     string   `json:"file_name"`
	QueryContent []string `json:"query_content"`
	messages     *ResponseMessages
	resources    *ResourceController
}

// handle the request
func (i *InsertPayLoad) HandleInsertion(w http.ResponseWriter, r *http.Request) {
	payload := new(InsertPayLoad)
	i.messages = getInstance()
	i.resources = getResourceControllerInstance()
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json, err := i.messages.MarshalErrMessage(0)
		if err != nil {
			writeServerErrorJSONResponse(w, json)
		} else {
			writeClientErrorJSONResponse(w, json)
		}
	} else {
		defer r.Body.Close()
		reqBody, _ := io.ReadAll(r.Body)
		json.Unmarshal(reqBody, &payload)
		err := i.resources.RemoteCSVFile(payload.User, payload.FileName, payload.QueryType, payload.QueryContent) // TODO - Marshal Error Messages
	}
}

// send a negative response to the client if the are:
// - directory problems, like invalid name/path or fileNotFoundException
// - Insetion problems.
// @param json encoded message passed by the caller
func writeServerErrorJSONResponse(w http.ResponseWriter, jsonMessage []byte) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonMessage)
}

// send a negative response to client due to Client error
// - Invalid method
// @param json encoded message passed by the caller
func writeClientErrorJSONResponse(w http.ResponseWriter, jsonMessage []byte) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonMessage)
}

// send a success response to client due to a corrent insertion
// in the csv files
// @param json encoded message passed by the caller
func writeSuccessJSONResponse(w http.ResponseWriter, jsonMessage []byte) {
	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonMessage)
}
