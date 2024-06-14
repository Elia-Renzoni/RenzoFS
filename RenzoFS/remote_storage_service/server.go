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

func init() {
	logger := api.RenzoFSCustomLogger{}
	logger.SetUpRenzoFSCustomLogger()
}

func main() {
	insertion := &api.InsertPayLoad{}
	read := &api.ReadPayLoad{}
	delete := &api.DeletePayLoad{}
	update := &api.UpdatePayLoad{}
	deleteDir := &api.DeleteDirPayLoad{}
	createDir := &api.CreateDirPayLoad{}
	fileInfo := &api.FileInfo{}
	stats := &api.StatsPayload{}
	deleteFile := &api.DeleteFilePayLoad{}

	handle := http.NewServeMux()
	handle.HandleFunc("/insert", insertion.HandleInsertion)
	handle.HandleFunc("/read/", read.HandleRead)
	handle.HandleFunc("/delete/", delete.HandleDelete)
	handle.HandleFunc("/update", update.HandleUpdate)
	handle.HandleFunc("/deletedir/", deleteDir.HandleDirElimination)
	handle.HandleFunc("/createdir", createDir.HandleDirCreation)
	handle.HandleFunc("/fileinfo/", fileInfo.HandleFileInfo)
	handle.HandleFunc("/stats/", stats.HandleStats)
	handle.HandleFunc("/deletefile/", deleteFile.HandleFileElimination)

	http.ListenAndServe(":8080", handle)
}
