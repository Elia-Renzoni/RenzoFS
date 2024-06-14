/**
*	@author Elia Renzoni
*	@date 20/02/2024
*	@brief Directory Creation Handler
**/

package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CreateDirPayLoad struct {
	DirToCreate string `json:"dir_to_create"`
	ResourceController
	ResponseMessages
	RenzoFSCustomLogger
}

func (c *CreateDirPayLoad) HandleDirCreation(w http.ResponseWriter, r *http.Request) {
	c.OpenLogFile()

	if r.Method != http.MethodPost {
		json, _ := c.MarshalErrMessage("Method Not Allowed")
		handleCreateDirResponse(w, methodNotAllowed, json)
	}
	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, c)
	fmt.Printf("%v", c.DirToCreate)
	if err := c.CreateNewDir(c.DirToCreate); err != nil {
		json, err := c.MarshalErrMessage(err.Error())
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			handleCreateDirResponse(w, serverError, json)
		}
	} else {
		json, err := c.Marshalsuccess(c.DirToCreate + " has been created")
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			handleCreateDirResponse(w, clientSucces, json)
		}
		c.WriteInLogFile(http.MethodPost + "\t" + c.DirToCreate)
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
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMessage)
	}
}
