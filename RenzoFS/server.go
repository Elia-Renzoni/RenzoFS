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
	read := &api.ReadPayLoad{}
	delete := &api.DeletePayLoad{}
	update := &api.UpdatePayLoad{}
	deleteDir := &api.DeleteDirPayLoad{}
	createDir := &api.CreateDirPayLoad{}

	handle := http.NewServeMux()
	handle.HandleFunc("/insert", insertion.HandleInsertion)
	handle.HandleFunc("/read", read.HandleRead)
	handle.HandleFunc("/delete", delete.HandleDelete)
	handle.HandleFunc("/update", update.HandleUpdate)
	handle.HandleFunc("/deletedir/{dir}/", deleteDir.HandleDirElimination)
	handle.HandleFunc("/createdir", createDir.HandleDirCreation)

	http.ListenAndServe(":8080", handle)
}
