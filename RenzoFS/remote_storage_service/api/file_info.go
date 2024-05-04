/**
*	@author Elia Renzoni
*	@date 22/03/2024
*	@brief File information handler
*
**/

package api

import (
	"net/http"
	"strconv"
	"strings"
)

type FileInfo struct {
	fileName   string
	dirName    string
	controller *ResourceController
	messages   *ResponseMessages
	logger     *RenzoFSCustomLogger
}

func (f *FileInfo) HandleFileInfo(w http.ResponseWriter, r *http.Request) {
	tmp := r.URL.Path                   // fileinfo/user/filename
	tmpSlice := strings.Split(tmp, "/") // [fileinfo, user, filename]
	f.fileName = tmpSlice[3]            // filename
	f.dirName = tmpSlice[2]             // dirname
	f.controller = getResourceControllerInstance()
	f.messages = getInstance()

	f.logger = GetRenzoFSCustomLogger()
	f.logger.OpenLogFile()

	if r.Method != http.MethodGet {
		json, err := f.messages.MarshalErrMessage("Method Not Valid")
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			handleFileInfoResponse(w, methodNotAllowed, json)
		}
	} else {
		fileInfo, err := f.controller.GetFileInformations(f.dirName, f.fileName)
		if err != nil {
			json, err := f.messages.MarshalErrMessage(err.Error())
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				handleFileInfoResponse(w, serverError, json)
			}
		} else {
			//var informations string = " " + fileInfo.Name() + " " + string(fileInfo.Size()) + " " + fileInfo.ModTime().String()
			var messageInfo [3]string = [3]string{
				fileInfo.Name(),
				strconv.Itoa(int(fileInfo.Size())),
				fileInfo.ModTime().GoString(),
			}
			json, err := f.messages.MarshalSuccessFileInformations(messageInfo)
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				handleFileInfoResponse(w, clientSucces, json)
				f.logger.WriteInLogFile("Readed file information from " + f.dirName + "/" + f.fileName + " in RenzoFS")
			}
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
