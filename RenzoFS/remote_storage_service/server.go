/**
*	@author Elia Renzoni
*	@date 01/03/2024
*	@version 1.0
*	@brief RenzoFS - Distributed File System
**/

package main

import (
	"net/http"
	"renzofs/remote_storage_service/api"
)

func main() {
	insertion := &api.InsertPayLoad{}
	read := &api.ReadPayLoad{}
	delete := &api.DeletePayLoad{}
	update := &api.UpdatePayLoad{}
	deleteDir := &api.DeleteDirPayLoad{}
	createDir := &api.CreateDirPayLoad{}
	fileInfo := &api.FileInfo{}

	handle := http.NewServeMux()
	handle.HandleFunc("/insert", insertion.HandleInsertion)
	handle.HandleFunc("/read/", read.HandleRead)
	handle.HandleFunc("/delete/", delete.HandleDelete)
	handle.HandleFunc("/update", update.HandleUpdate)
	handle.HandleFunc("/deletedir/", deleteDir.HandleDirElimination)
	handle.HandleFunc("/createdir", createDir.HandleDirCreation)
	handle.HandleFunc("/fileinfo/", fileInfo.HandleFileInfo)

	http.ListenAndServe(":8080", handle)
}
