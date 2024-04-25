/**
*	@author Elia Renzoni
*	@date 20/02/2024
*	@brief Tipical Errors and Success situation
**/

package api

const (
	methodNotAllowed byte = iota
	serverError           // impossible to complete the request
	clientSucces          // request get succesful
)
