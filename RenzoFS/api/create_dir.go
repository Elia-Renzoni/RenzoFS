package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type CreateDirPayLoad struct {
	DirToCreate string `json:"dir_to_create"`
	controller  *ResourceController
	messages    *ResponseMessages
}

func (c *CreateDirPayLoad) HandleDirCreation(w http.ResponseWriter, r *http.Request) {
	payload := new(CreateDirPayLoad)
	c.messages = getInstance()
	c.controller = getResourceControllerInstance()
	if r.Method != http.MethodPost {
		json, _ := c.messages.MarshalErrMessage("Method Not Allowed")
		handleCreateDirResponse(w, methodNotAllowed, json)
	}
	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, payload)
	if err := c.controller.CreateNewDir(payload.DirToCreate); err != nil {
		json, _ := c.messages.MarshalErrMessage(err.Error())
		handleCreateDirResponse(w, serverError, json)
	} else {
		json, _ := c.messages.Marshalsuccess(payload.DirToCreate + " has been created")
		handleCreateDirResponse(w, clientSucces, json)
	}
}

func handleCreateDirResponse(w http.ResponseWriter, id byte, jsonMessage []byte) {
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
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMessage)
	}
}
