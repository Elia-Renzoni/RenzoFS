/**
*	@author Elia Renzoni
*	@date 02/03/2024
*	@brief Check if the directory and files exist
*
**/

package api

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"path/filepath"
)

type ResourceController struct {
}

var resourceControllerInstace *ResourceController

// crud operations
const (
	insert string = "insert"
	update string = "update"
	delete string = "delete"
	read   string = "read"
)

// singleton pattern
func getResourceControllerInstance() *ResourceController {
	if resourceControllerInstace == nil {
		return new(ResourceController)
	}
	return resourceControllerInstace
}

func (r *ResourceController) DreateNewDir(dirname string) {
	if err := os.Mkdir(filepath.Join("local_file_system", dirname), os.ModeDir); err != nil {
		log.Fatal(err)
	}
}

func (r *ResourceController) DeleteDir(dirname string) {
	
}

// function to call when the client need to
// - READ
// - DELETE
// - INSERT
// - UPDATE
// information in files
// @param dir of the local file system - username
// @param file name
// @param id of the CRUD operation
// @param query content
func (r *ResourceController) RemoteCSVFile(dir string, filename string, queryType string, query []string) error {
	var err error
	switch queryType {
	case insert:
		err = writeRemoteCSV(dir, filename, query)
	case read:
		err = readInRemoteCSV(dir, filename, query)
	case update:
		err = updateRemoteCSV(dir, filename, query)
	case delete:
		err = deleteRemoteCSV(dir, filename, query)
	default:
		err = errors.New("Invalid Query Type")
	}
	return err
}

func writeRemoteCSV(dir, filename string, query []string) error {
	var changer func(string) error = changeWorkDirectory
	var switchToNormal func() error = switchToRenzoFSDir
	
	if err := changer(dir); err != nil {
		return err
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close() // defer-close pattern
	writer := csv.NewWriter(file)
	defer writer.Flush()
	if err := writer.Write(query); err != nil {
		return err
	}
	defer switchToNormal()
	return nil
}

// TODO
func readInRemoteCSV(dir, filename string, query []string) error {
	var changer func(string) error = changeWorkDirectory
	var switchToNormal func() error = switchToRenzoFSDir

	if err := changer(dir); err != nil {
		return err
	}
	
	file, err := os.Open(filename, os.O_RONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close();
	reader := csv.NewReader(file)
	// reader logic
	defer switchToNormal()
	return nil
}

// TODO
func updateRemoteCSV(dir, filename string, query []string) error {
	return nil
}

// TODO
func deleteRemoteCSV(dir, filename string, query []string) error {
	return nil
}

func changeWorkDirectory(dir string) error {
	for {
		if err := os.Chdir(filepath.Join("local_file_system", dir)); err != nil {
			return err
		} else {
			break;
		}
	}
	return nil
}

func switchToRenzoFSDir() error {
	if err := os.Chdir("E:/RenzoFS"); err != nil {
		return err 
	}
	return nil
}
