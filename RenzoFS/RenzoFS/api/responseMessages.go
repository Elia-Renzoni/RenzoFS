/**
*	@author Elia Renzoni
*	@date 02/02/2024
*	@brief Response Marshaller
**/

package api

import "encoding/json"

const (
	methodNotAllowed string = "The Method is not allowed for the specific call"
	invalidParameter string = "Invalid query parameter, check again and send back"
	okMessage        string = "Good Request, Good Parameters"
)

type ResponseMessages struct {
	ErrMessage map[string]string
	OkMessage  map[string]string
}

func (r *ResponseMessages) MarshallErrMessage() (jsonErrMessage []byte, err error) {
	r.ErrMessage = make(map[string]string)
	r.ErrMessage["err_message"] = methodNotAllowed
	jsonErrMessage, err = json.Marshal(r.ErrMessage)
	return
}

func (r *ResponseMessages) MarshallOkMessage() (jsonOkMessage []byte, err error) {
	r.OkMessage = make(map[string]string)
	r.OkMessage["ok_message"] = okMessage
	jsonOkMessage, err = json.Marshal(r.OkMessage)
	return
}
