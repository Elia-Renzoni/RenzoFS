/**
*	@author Elia Renzoni
*	@date 02/02/2024
*	@brief Response Marshaller
**/

package api

import "encoding/json"

// error and exceptions
const (
	methodNotAllowed string = "The Method is not allowed for the specific call"
	invalidParameter string = "Invalid query parameter, check again and send back"
	okMessage        string = "Good Request, Good Parameters"
	dirException     string = "Directory not exitst, create a new one"
	fileException    string = "File not exist, create a new one"
	POSTokMessage    string = "Added query content"
)

const (
	errorMessage, fineMessage string = "err_message", "ok_message"
)

type ResponseMessages struct {
	ErrMessage map[string]string
	OkMessage  map[string]string
}

func (r *ResponseMessages) MarshallErrMessage() (jsonErrMessage []byte, err error) {
	r.ErrMessage = make(map[string]string)
	r.ErrMessage[errorMessage] = methodNotAllowed
	jsonErrMessage, err = json.Marshal(r.ErrMessage)
	return
}

func (r *ResponseMessages) MarshallOkMessage() (jsonOkMessage []byte, err error) {
	r.OkMessage = make(map[string]string)
	r.OkMessage[fineMessage] = okMessage
	jsonOkMessage, err = json.Marshal(r.OkMessage)
	return
}

func (r *ResponseMessages) MarshallDirException() (jsonErrMessage []byte, err error) {
	r.ErrMessage = make(map[string]string)
	r.ErrMessage[errorMessage] = dirException
	jsonErrMessage, err = json.Marshal(r.ErrMessage)
	return
}

func (r *ResponseMessages) MarshallFileException() (jsonErrMessage []byte, err error) {
	r.ErrMessage = make(map[string]string)
	r.ErrMessage[errorMessage] = fileException
	jsonErrMessage, err = json.Marshal(r.ErrMessage)
	return
}

func (r *ResponseMessages) MarshallPOSTSuccess() (jsonPOSTSuccessMessage []byte, err error) {
	r.OkMessage = make(map[string]string)
	r.OkMessage[okMessage] = POSTokMessage
	jsonPOSTSuccessMessage, err = json.Marshal(r.OkMessage)
	return
}
