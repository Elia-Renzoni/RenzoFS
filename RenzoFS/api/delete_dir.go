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
}

func (d *DeleteDirPayLoad) HandleDirElimination(w http.ResponseWriter, r *http.Request) {
	tmp := r.URL.Path               // /deletedir/dirname
	tmp2 := strings.Split(tmp, "/") // [deletedir, dirname]
	d.dirToDelete = tmp2[2]         // dirname
	fmt.Printf("%v", d.dirToDelete)
	d.controller = getResourceControllerInstance()
	d.messages = getInstance()
	if r.Method != http.MethodDelete {
		json, _ := d.messages.MarshalErrMessage("Method Not Allowed")
		handleDeleteDirResponse(w, methodNotAllowed, json)
	} else {
		if err := d.controller.DeleteDir(d.dirToDelete); err != nil {
			json, _ := d.messages.MarshalErrMessage(err.Error())
			handleDeleteDirResponse(w, serverError, json)
		} else {
			json, _ := d.messages.Marshalsuccess(d.dirToDelete + " Has Been Deleted")
			handleDeleteDirResponse(w, clientSucces, json)
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
