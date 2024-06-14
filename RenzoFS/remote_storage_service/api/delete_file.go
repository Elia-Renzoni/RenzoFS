package api

import (
	"net/http"
	"path"
	"strings"
)

type DeleteFilePayLoad struct {
	dirName, fileName string
	ResourceController
	ResponseMessages
	RenzoFSCustomLogger
}

func (d *DeleteFilePayLoad) HandleFileElimination(w http.ResponseWriter, r *http.Request) {
	parameters := strings.Split(r.URL.Path, "/")
	d.dirName = parameters[1]
	d.fileName = parameters[2]

	if r.Method != http.MethodDelete {
		json, err := d.MarshalErrMessage("Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		handleFileEliminationResponses(w, methodNotAllowed, json)
	} else {
		if err := d.DeleteFile(path.Join()); err != nil {
			json, err := d.MarshalErrMessage(err.Error())
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				handleFileEliminationResponses(w, serverError, json)
			}
		} else {
			json, err := d.Marshalsuccess(d.fileName + " Has Been Deleted")
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				handleFileEliminationResponses(w, clientSucces, json)
			}
			d.WriteInLogFile(http.MethodDelete + "\t" + d.fileName)
		}
	}
}

func handleFileEliminationResponses(w http.ResponseWriter, id byte, jsonMessage []byte) {
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
