/**
*	@author Elia Renzoni
*	@date 02/02/2024
*	@brief Response Marshaller
**/

package api

import (
	"encoding/json"
	"errors"
)

const (
	errorMessage, succMessage string = "err_message", "success_message"
)

// errot types enumeration
const (
	methodNotAllowed byte = iota
	invalidParameter
	SuccessMsg
	dirException
	fileException
	POSTSuccessMesssage
	invalidMessageType
)

// K-V store
// K = message id
// V = message content
var responseSettings map[byte]string = map[byte]string{
	methodNotAllowed:    "The Method is not allowed for the specific call",
	invalidParameter:    "Invalid query parameter, check again and send back",
	SuccessMsg:          "Good Request, Good Parameters",
	dirException:        "Directory not exitst, create a new one",
	fileException:       "File not exist, create a new one",
	POSTSuccessMesssage: "Added query content",
	invalidMessageType:  "Invalid Error Type",
}

type ResponseMessages struct {
	errMessage      map[string]string
	succcessMessage map[string]string
}

// global variables used to check the memory assignement
// used for singleton pattern
var instance *ResponseMessages

// singleton pattern
func getInstance() *ResponseMessages {
	if instance == nil {
		return new(ResponseMessages)
	}
	return instance
}

// marshal errors
func (r *ResponseMessages) MarshalErrMessage(messageType byte) (jsonErrMessage []byte, err error) {
	r.errMessage = make(map[string]string)
	messageToMarshal, ok := responseSettings[messageType] // comma ok pattern
	if !ok {
		message := responseSettings[6]
		return []byte(message), errors.New("Invalid Message Type")
	}
	r.errMessage[errorMessage] = messageToMarshal
	jsonErrMessage, err = json.Marshal(r.errMessage)
	return
}

// marshal success responses
func (r *ResponseMessages) Marshalsuccess(messageType byte) (jsonSuccessMessage []byte, err error) {
	r.succcessMessage = make(map[string]string)
	messageToMarshal, ok := responseSettings[messageType] // comma ok pattern
	if !ok {
		message := responseSettings[6]
		return []byte(message), errors.New("Invalid Message Type")
	}
	r.succcessMessage[succMessage] = messageToMarshal
	jsonSuccessMessage, err = json.Marshal(r.succcessMessage)
	return
}
