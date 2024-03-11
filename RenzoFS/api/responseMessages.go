/**
*	@author Elia Renzoni
*	@date 02/02/2024
*	@brief Response Marshaller
**/

package api

import (
	"encoding/json"
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
)

var responseSettings map[byte]string = map[byte]string{
	methodNotAllowed:    "The Method is not allowed for the specific call",
	invalidParameter:    "Invalid query parameter, check again and send back",
	SuccessMsg:          "Good Request, Good Parameters",
	dirException:        "Directory not exitst, create a new one",
	fileException:       "File not exist, create a new one",
	POSTSuccessMesssage: "Added query content",
}

type ResponseMessages struct {
	ErrMessage      map[string]string
	SucccessMessage map[string]string
}

func (r *ResponseMessages) MarshalErrMessage(messageType byte) (jsonErrMessage []byte, err error) {
	r.ErrMessage = make(map[string]string)
	messageToMarshal := getMessageToMarshal(messageType)
	r.ErrMessage[errorMessage] = messageToMarshal
	if messageToMarshal == "" {
		panic("Invalid Message Type")
	}
	jsonErrMessage, err = json.Marshal(r.ErrMessage)
	return
}

func (r *ResponseMessages) MarshalSuccessMessage(messageType byte) (jsonSuccessMessage []byte, err error) {
	r.SucccessMessage = make(map[string]string)
	messageToMarshal := getMessageToMarshal(messageType)
	r.SucccessMessage[succMessage] = messageToMarshal
	if messageToMarshal == "" {
		panic("Invalid Message Type")
	}
	jsonSuccessMessage, err = json.Marshal(r.SucccessMessage)
	return
}

func getMessageToMarshal(messageType byte) (message string) {
	for key, value := range responseSettings {
		if key == messageType {
			return value
		}
	}
	return ""
}
