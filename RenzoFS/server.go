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
	insertion := &api.InsertPayLoad{}
	handle := http.NewServeMux()
	handle.HandleFunc("/insert", insertion.HandleInsertion)
	/*handle.HandleFunc("/read", api.HandleRead)
	handle.HandleFunc("/update", api.HandleUpdate)
	handle.HandleFunc("/delete", api.HandleDelete)
	handle.HandleFunc("/delete/dir", api.HandleDirElimination)
	handle.HandleFunc("/create/dir", api.HandleDirCreation)*/

	http.ListenAndServe(":8080", handle)
}
