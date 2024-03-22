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
	i.messages = getInstance()
	i.resources = getResourceControllerInstance()
	if r.Method != http.MethodPost {
		json, _ := i.messages.MarshalErrMessage("Method Not Allowed")
		handleInsertResponse(w, methodNotAllowed, json)
	} else {
		// read the request
		defer r.Body.Close()
		reqBody, _ := io.ReadAll(r.Body)
		json.Unmarshal(reqBody, i)
		err := i.resources.RemoteCSVFile(i.User, i.FileName, i.QueryType, i.QueryContent) // TODO - Marshal Error Messages
		if err != nil {
			jsonMessage, _ := i.messages.MarshalErrMessage(err.Error())
			handleInsertResponse(w, serverError, jsonMessage)
		} else {
			jsonMessage, _ := i.messages.Marshalsuccess("information successfully added")
			handleInsertResponse(w, clientSucces, jsonMessage)
		}
	}
}

func handleInsertResponse(w http.ResponseWriter, id byte, jsonMessage []byte) {
	switch id {
	case methodNotAllowed:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMessage)
	case serverError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMessage)
	case clientSucces:
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMessage)
	}
}
