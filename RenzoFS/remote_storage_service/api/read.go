/**
*	@author Elia Renzoni
*	@brief this module handles the GET requests made to the server,
*	in particular it returns to the client the information associated
*	with the id indicated in the query
**/

package api

import (
	"net/http"
	"net/url"
	"strings"
)

type ReadPayLoad struct {
	// directory name and inner file name
	user, fileName string

	// URL written by the client
	url url.Values // map[string][]string

	// composition fields
	controller *ResourceController
	messages   *ResponseMessages

	logger *RenzoFSCustomLogger
}

// this method handle the GET crud operation
// access to the specified file and return
// the content that match with the written id
func (r *ReadPayLoad) HandleRead(w http.ResponseWriter, req *http.Request) {
	tmp := req.URL.Path             // read/user/filename
	tmp2 := strings.Split(tmp, "/") // [read, user, filename]
	r.user = tmp2[2]                // user
	r.fileName = tmp2[3]            // filename
	r.url = req.URL.Query()         // get the url query section
	r.controller = getResourceControllerInstance()
	r.messages = getInstance()

	r.logger = GetRenzoFSCustomLogger()
	r.logger.OpenLogFile()

	if req.Method != http.MethodGet {
		json, err := r.messages.MarshalErrMessage("Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		handleGetResponses(w, methodNotAllowed, json)
	} else {
		responseToEncode, err := r.controller.ReadInRemoteCSV(r.user, r.fileName, "read", r.url)
		if err != nil {
			json, err := r.messages.MarshalErrMessage(err.Error())
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
			r.logger.WriteInLogFile(http.MethodGet + "\t" + r.user + "\t" + r.fileName)
		}
	}
}

// this function set the responses that the
// server has to write to the client
// @param id indicates the type of response @see enumeration
// @param jsonMessage indicates the JSON response to write in
// the response payload
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
