/**
*	@author Elia Renzoni
*	@date 22/04/2024
*	@brief This module handles the insertions of information in files by users
**/

package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type InsertPayLoad struct {
	// contain the type of the query
	// in this case it must be an insert
	// query
	QueryType string `json:"query_type"`

	// contain the directory name
	User string `json:"user"`

	// contain the file name
	FileName string `json:"file_name"`

	// contain the query content that is a JSON array
	QueryContent []string `json:"query_content"`

	// composition fields
	ResponseMessages
	ResourceController
	RenzoFSCustomLogger
}

// this method handle the POST operation thath is
// related to the insertion of new information to
// the specified file
func (i *InsertPayLoad) HandleInsertion(w http.ResponseWriter, r *http.Request) {
	i.OpenLogFile()

	if r.Method != http.MethodPost {
		json, err := i.MarshalErrMessage("Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			handleInsertResponse(w, methodNotAllowed, json)
		}
	} else {
		// read the request
		defer r.Body.Close()
		reqBody, _ := io.ReadAll(r.Body)
		json.Unmarshal(reqBody, i)
		err := i.WriteRemoteCSV(i.User, i.FileName, i.QueryType, i.QueryContent)
		if err != nil {
			jsonMessage, err := i.MarshalErrMessage(err.Error())
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				handleInsertResponse(w, serverError, jsonMessage)
			}
		} else {
			jsonMessage, err := i.Marshalsuccess("information successfully added")
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				handleInsertResponse(w, clientSucces, jsonMessage)
				i.WriteInLogFile(http.MethodPost + "\t" + i.User + "\t" + i.FileName)
			}
		}
	}
}

// this function set the responses that the
// server has to write to the client
// @param id indicate the type of response @see enumeration
// @param jsonMessage indicate the JSON response to write in
// te response payload
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
