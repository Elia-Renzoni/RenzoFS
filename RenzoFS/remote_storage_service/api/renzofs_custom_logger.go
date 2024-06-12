/**
*	@author Elia Renzoni
*	@date 04/05/2024
*	@brief Remote Storage Service Custom Logger
*
**/

package api

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
)

type RenzoFSCustomLogger struct {
	InfoLogger *log.Logger
	// TODO
	mutex sync.Mutex
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
	return log.New(file, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func (l *RenzoFSCustomLogger) OpenLogFile() {
	os.OpenFile("renzo_fs_log.txt", os.O_APPEND|os.O_RDWR, 0666)
}

func (l *RenzoFSCustomLogger) WriteInLogFile(message string) {
	l.InfoLogger.Println(message)
}

func (l *RenzoFSCustomLogger) SearchInLogFile(dir, filename string) ([]string, error) {
	var response []string = make([]string, 0)
	file, err := os.Open("renzo_fs_log.txt")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read lines
	for scanner.Scan() {
		splittedContent := strings.Split(scanner.Text(), "\t")
		for index := range splittedContent {
			if splittedContent[index] == dir || splittedContent[index] == filename {
				response = append(response, scanner.Text())
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return response, nil
}
