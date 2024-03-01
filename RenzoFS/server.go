/**
*	@author Elia Renzoni
*	@date 01/03/2024
*	@version 1.0
*	@brief RenzoFS - Distributed File System
**/

package main

import (
	"net/http"
	"renzofs/api"
)

func main() {

	handle := http.NewServeMux()

	handle.HandleFunc("/insert", api.HandleInsertion)
	// TODO - other handlers

	http.ListenAndServe(":8080", nil)
}
