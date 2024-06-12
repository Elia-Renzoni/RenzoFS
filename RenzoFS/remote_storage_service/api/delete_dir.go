/**
*	@author Elia Renzoni
*	@date 20/03/2024
*	@brief Delete Directory Endpoint Handler
**/

package api

import (
	"fmt"
	"net/http"
	"strings"
)

type DeleteDirPayLoad struct {
	dirToDelete string
	controller  *ResourceController
	messages    *ResponseMessages
	logger      *RenzoFSCustomLogger
}

func (d *DeleteDirPayLoad) HandleDirElimination(w http.ResponseWriter, r *http.Request) {
	tmp := r.URL.Path               // /deletedir/dirname
	tmp2 := strings.Split(tmp, "/") // [deletedir, dirname]
	d.dirToDelete = tmp2[2]         // dirname
	fmt.Printf("%v", d.dirToDelete)
	d.controller = getResourceControllerInstance()
	d.messages = getInstance()

	d.logger = GetRenzoFSCustomLogger()
	d.logger.OpenLogFile()

	if r.Method != http.MethodDelete {
		json, err := d.messages.MarshalErrMessage("Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			handleDeleteDirResponse(w, methodNotAllowed, json)
		}
	} else {
		if err := d.controller.DeleteDir(d.dirToDelete); err != nil {
			json, err := d.messages.MarshalErrMessage(err.Error())
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				handleDeleteDirResponse(w, serverError, json)
			}
		} else {
			json, err := d.messages.Marshalsuccess(d.dirToDelete + " Has Been Deleted")
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				handleDeleteDirResponse(w, clientSucces, json)
			}
			d.logger.WriteInLogFile(http.MethodDelete + "\t" + d.dirToDelete)
		}
	}
}

func handleDeleteDirResponse(w http.ResponseWriter, id byte, jsonMessage []byte) {
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
