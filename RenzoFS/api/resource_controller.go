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

func (r *ResourceController) CreateNewDir(dirname string) error {
	if err := os.Mkdir(filepath.Join("local_file_system", dirname), os.ModeDir); err != nil {
		return err
	}
	return nil
}

func (r *ResourceController) DeleteDir(dirname string) error {
	defer func() error {
		if err := os.Chdir("E:/RenzoFS"); err != nil {
			return err
		}
		return nil
	}()

	for {
		// change work directory to local_file_system + user dir
		if err := os.Chdir(filepath.Join("local_file_system")); err != nil {
			return err
		} else {
			break
		}
	}
	if err := os.Remove(dirname); err != nil {
		return err
	}

	return nil
}

// TODO
func (r *ResourceController) GetFileInformations(dirname, filename string) (os.FileInfo, error) {
	var (
		fileInfo os.FileInfo
		err      error
	)
	for {
		// change work directory to local_file_system + user dir
		if err := os.Chdir(filepath.Join("local_file_system", dirname)); err != nil {
			return fileInfo, err
		} else {
			break
		}
	}
	fileInfo, err = os.Stat(filename)
	if err != nil {
		return fileInfo, err
	}
	return fileInfo, nil
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
	for {
		// change work directory to local_file_system + user dir
		if err := os.Chdir(filepath.Join("local_file_system", dir)); err != nil {
			return err
		} else {
			break
		}
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

	// change work directory
	defer func() error {
		if err := os.Chdir("E:/RenzoFS"); err != nil {
			return err
		}
		return nil
	}()

	return nil
}

// TODO
func readInRemoteCSV(dir, filename string, query []string) error {
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
