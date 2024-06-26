/**
*	@author Elia Renzoni
*	@date 02/02/2024
*	@brief Response Marshaler
**/

package api

import (
	"encoding/json"
)

const (
	errorMessage, succMessage   string = "err_message", "success_message"
	fileName, fileSize, modTime string = "name", "size", "modification time"
)

type ResponseMessages struct {
	errMessage      map[string]string
	succcessMessage map[string]string
}

// marshal errors
func (r *ResponseMessages) MarshalErrMessage(messageText string) (jsonErrMessage []byte, err error) {
	r.errMessage = make(map[string]string)
	r.errMessage[errorMessage] = messageText
	jsonErrMessage, err = json.Marshal(r.errMessage)
	return
}

// marshal success responses
func (r *ResponseMessages) Marshalsuccess(messageText string) (jsonSuccessMessage []byte, err error) {
	r.succcessMessage = make(map[string]string)
	r.succcessMessage[succMessage] = messageText
	jsonSuccessMessage, err = json.Marshal(r.succcessMessage)
	return
}

// marhsal success response in case of file information get request
func (r *ResponseMessages) MarshalSuccessFileInformations(messages [3]string) (jsonSuccessMessage []byte, err error) {
	r.succcessMessage = make(map[string]string)
	r.succcessMessage[fileName] = messages[0]
	r.succcessMessage[fileSize] = messages[1]
	r.succcessMessage[modTime] = messages[2]
	jsonSuccessMessage, err = json.Marshal(r.succcessMessage)
	return
}

func (r *ResponseMessages) MarshalSuccesReadResults(messageToEncode map[string]string) (jsonSuccessMessage []byte, err error) {
	jsonSuccessMessage, err = json.Marshal(messageToEncode)
	return
}

func (r *ResponseMessages) MarshalSuccesStatResults(messageToEncode []string) (jsonSuccessMessage []byte, err error) {
	r.succcessMessage = make(map[string]string)
	var toEncode string
	for index := range messageToEncode {
		toEncode += messageToEncode[index]
	}
	r.succcessMessage[succMessage] = toEncode
	jsonSuccessMessage, err = json.Marshal(r.succcessMessage)
	return
}
