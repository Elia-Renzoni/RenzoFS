package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type UpdatePayLoad struct {
	QueryType    string              `json:"query_type"`
	User         string              `json:"user_name"`
	FileName     string              `json:"file_name"`
	QueryContent map[string][]string `json:"query_content"`
	ResourceController
	ResponseMessages
	RenzoFSCustomLogger
}

func (u *UpdatePayLoad) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	u.OpenLogFile()

	if r.Method != http.MethodPatch {
		json, err := u.MarshalErrMessage("Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		handleUpdateResponses(w, methodNotAllowed, json)
	} else {
		defer r.Body.Close()
		reqBody, _ := io.ReadAll(r.Body)
		json.Unmarshal(reqBody, u)
		err := u.UpdateRemoteCSV(u.User, u.FileName, u.QueryType, u.QueryContent)
		if err != nil {
			json, err := u.MarshalErrMessage(err.Error())
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				handleUpdateResponses(w, serverError, json)
			}
		} else {
			json, err := u.Marshalsuccess("Informations Succesfully Updated")
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				handleUpdateResponses(w, clientSucces, json)
				u.WriteInLogFile(http.MethodPatch + " " + u.User + " " + u.FileName)
			}
		}
	}
}

func handleUpdateResponses(w http.ResponseWriter, id byte, jsonMessage []byte) {
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
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMessage)
	}
}
