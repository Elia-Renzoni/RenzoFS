/**
*	@author Elia Renzoni
*	@date 22/03/2024
*	@brief File information handler
*
**/

package api

import (
	"net/http"
	"strings"
)

type FileInfo struct {
	fileName   string
	dirName    string
	controller *ResourceController
	messages   *ResponseMessages
}

func (f *FileInfo) HandleFileInfo(w http.ResponseWriter, r *http.Request) {
	tmp := r.URL.Path                   // fileinfo/user/filename
	tmpSlice := strings.Split(tmp, "/") // [fileinfo, user, filename]
	f.fileName = tmpSlice[3]            // filename
	f.dirName = tmpSlice[2]             // dirname
	f.controller = getResourceControllerInstance()
	f.messages = getInstance()
	if r.Method != http.MethodGet {
		json, _ := f.messages.MarshalErrMessage("Method Not Valid")
		handleFileInfoResponse(w, methodNotAllowed, json)
	} else {
		fileInfo, err := f.controller.GetFileInformations(f.dirName, f.fileName)
		if err != nil {
			json, _ := f.messages.MarshalErrMessage(err.Error())
			handleFileInfoResponse(w, serverError, json)
		} else {
			var informations string = " " + fileInfo.Name() + " " + string(fileInfo.Size()) + " " + fileInfo.ModTime().String()
			json, _ := f.messages.Marshalsuccess(informations)
			handleFileInfoResponse(w, clientSucces, json)
		}
	}
}

func handleFileInfoResponse(w http.ResponseWriter, id byte, jsonMessage []byte) {
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
