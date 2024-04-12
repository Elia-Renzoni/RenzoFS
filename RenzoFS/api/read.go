package api

import (
	"net/http"
	"net/url"
	"strings"
)

type ReadPayLoad struct {
	user, fileName string
	url            url.Values // map[string][]string
	controller     *ResourceController
	messages       *ResponseMessages
}

func (r *ReadPayLoad) HandleRead(w http.ResponseWriter, req *http.Request) {
	tmp := req.URL.Path             // read/user/filename
	tmp2 := strings.Split(tmp, "/") // [read, user, filename]
	r.user = tmp2[1]                // user
	r.fileName = tmp2[2]            // filename
	r.url = req.URL.Query()
	r.controller = getResourceControllerInstance()
	r.messages = getInstance()
	if req.Method != http.MethodGet {
		json, err := r.messages.MarshalErrMessage("Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		handleGetResponses(w, methodNotAllowed, json)
	} else {
		responseToEncode, err := r.controller.ReadInRemoteCSV(r.user, r.fileName, "read", r.url)
		if err != nil {
			json, err := r.messages.MarshalErrMessage("Internal Server Error")
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
			handleGetResponses(w, serverError, json)
		} else {
			json, err := r.messages.MarshalSuccesReadResults(responseToEncode)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
			handleGetResponses(w, clientSucces, json)
		}
	}
}

func handleGetResponses(w http.ResponseWriter, id byte, jsonMessage []byte) {
	switch id {
	case serverError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMessage)
	case methodNotAllowed:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMessage)
	case clientSucces:
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMessage)
	}
}
