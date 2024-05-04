/**
*	@author Elia Renzoni
*	@date 04/05/2024
*	@brief Remote Storage Service Custom Logger
*
**/

package api

import (
	"log"
	"os"
)

type RenzoFSCustomLogger struct {
	InfoLogger *log.Logger
}

var customLoggerLocalInstance *RenzoFSCustomLogger

func newRenzoFSCustomLogger() *RenzoFSCustomLogger {
	return &RenzoFSCustomLogger{
		InfoLogger: setUpRenzoFSCustomLogger(),
	}
}

// singelton
func GetRenzoFSCustomLogger() *RenzoFSCustomLogger {
	if customLoggerLocalInstance == nil {
		customLoggerLocalInstance = newRenzoFSCustomLogger()
	}
	return customLoggerLocalInstance
}

func setUpRenzoFSCustomLogger() *log.Logger {
	var file, err = os.Create("renzo_fs_log.txt")
	if err != nil {
		panic(err)
	}
	return log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *RenzoFSCustomLogger) OpenLogFile() {
	os.OpenFile("renzo_fs_log.txt", os.O_APPEND|os.O_RDWR, 0666)
}

func (l *RenzoFSCustomLogger) WriteInLogFile(message string) {
	l.InfoLogger.Println(message)
}
