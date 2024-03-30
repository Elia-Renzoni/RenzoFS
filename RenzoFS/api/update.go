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
	controller   *ResourceController
	messages     *ResponseMessages
}

func (u *UpdatePayLoad) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	u.controller = getResourceControllerInstance()
	u.messages = getInstance()
	if r.Method != http.MethodPatch {
		json, err := u.messages.MarshalErrMessage("Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		handleUpdateResponses(w, methodNotAllowed, json)
	} else {
		defer r.Body.Close()
		reqBody, _ := io.ReadAll(r.Body)
		json.Unmarshal(reqBody, u)
		err := u.controller.UpdateRemoteCSV(u.User, u.FileName, u.QueryType, u.QueryContent)
		if err != nil {
			json, err := u.messages.MarshalErrMessage(err.Error())
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				handleUpdateResponses(w, serverError, json)
			}
		} else {
			json, err := u.messages.Marshalsuccess("Informations Succesfully Updated")
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				handleUpdateResponses(w, clientSucces, json)
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
