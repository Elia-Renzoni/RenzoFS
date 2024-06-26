package api

import (
	"fmt"
	"net/http"
	"strings"
)

type DeleteFilePayLoad struct {
	dirname, filename string
	ResourceController
	ResponseMessages
	RenzoFSCustomLogger
}

func (df *DeleteFilePayLoad) HandleFileElimination(w http.ResponseWriter, r *http.Request) {
	request := r.URL.Path
	splittedRequest := strings.Split(request, "/")
	df.dirname = splittedRequest[2]
	df.filename = splittedRequest[3]
	fmt.Printf("%v - %v", df.dirname, df.filename)

	if r.Method != http.MethodDelete {
		json, err := df.MarshalErrMessage("Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			handleFileEliminationResponses(w, methodNotAllowed, json)
		}
	} else {
		if err := df.DeleteFile(df.dirname, df.filename); err != nil {
			json, errs := df.MarshalErrMessage(err.Error())
			if errs != nil {
				http.Error(w, errs.Error(), 500)
			} else {
				handleFileEliminationResponses(w, serverError, json)
			}
		} else {
			json, err := df.Marshalsuccess("File Succesfully Deleted")
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				handleFileEliminationResponses(w, clientSucces, json)
			}
		}
	}
}

func handleFileEliminationResponses(w http.ResponseWriter, id byte, jsonMessage []byte) {
	switch id {
	case methodNotAllowed:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content/Type", "application/json")
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
