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
	"time"
)

func init() {
	logger := api.NewRenzoFSCustomLogger()
	logger.OpenLogFile()
}

func main() {
	insertion := &api.InsertPayLoad{}
	read := &api.ReadPayLoad{}
	delete := &api.DeletePayLoad{}
	update := &api.UpdatePayLoad{}
	deleteDir := &api.DeleteDirPayLoad{}
	createDir := &api.CreateDirPayLoad{}
	fileInfo := &api.FileInfo{}
	deletefile := &api.DeleteFilePayLoad{}
	stats := &api.StatsPayload{}
	health := &api.HealthSystems{}

	handle := http.NewServeMux()
	handle.HandleFunc("/insert", insertion.HandleInsertion)
	handle.HandleFunc("/read/", read.HandleRead)
	handle.HandleFunc("/delete/", delete.HandleDelete)
	handle.HandleFunc("/update", update.HandleUpdate)
	handle.HandleFunc("/deletedir/", deleteDir.HandleDirElimination)
	handle.HandleFunc("/createdir", createDir.HandleDirCreation)
	handle.HandleFunc("/fileinfo/", fileInfo.HandleFileInfo)
	handle.HandleFunc("/deletefile/", deletefile.HandleFileElimination)
	handle.HandleFunc("/stats/", stats.HandleStats)
	handle.HandleFunc("/health", health.HandleHealthCheck)

	remoteSServer := &http.Server{
		Addr:              ":8080",
		ReadTimeout:       3 * time.Second,
		WriteTimeout:      3 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           handle,
	}

	remoteSServer.ListenAndServe()
}
