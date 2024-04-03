package api

import (
	"net/http"
	"net/url"
	"strings"
)

type DeletePayLoad struct {
	user, fileName string
	url            url.Values
	controller     *ResourceController
	messages       *ResponseMessages
}

func (d *DeletePayLoad) HandleDelete(w http.ResponseWriter, r *http.Request) {
	d.controller = getResourceControllerInstance()
	d.messages = getInstance()
	tmp := r.URL.Path
	tmp2 := strings.Split(tmp, "/")
	d.user = tmp2[1]
	d.fileName = tmp2[2]
	d.url = r.URL.Query()
	if r.Method != http.MethodDelete {
		json, err := d.messages.MarshalErrMessage("Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		handleDeleteResponses(w, methodNotAllowed, json)
	} else {
		if err := d.controller.DeleteRemoteCSV(d.user, d.fileName, "delete", d.url); err != nil {
			json, err := d.messages.MarshalErrMessage(err.Error())
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
			handleDeleteResponses(w, serverError, json)
		} else {
			json, err := d.messages.Marshalsuccess("Informations succesfully eliminated")
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
			handleDeleteResponses(w, clientSucces, json)
		}
	}
}

func handleDeleteResponses(w http.ResponseWriter, id byte, jsonMessage []byte) {
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
