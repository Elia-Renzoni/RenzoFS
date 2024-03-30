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

// change the worker dir to path specified
type changeWorkDir func(string) error

// goes back to renzofs main dir
type backToHomeDir func() error

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

// *
func (r *ResourceController) DeleteDir(dirname string) error {
	var (
		firstDirChange changeWorkDir = changeWorkerDirectory
		lastDirChange  backToHomeDir = changeToMainDirectory
	)

	defer lastDirChange()

	// change to local_file_system
	if err := firstDirChange(""); err != nil {
		return err
	}

	if err := os.Remove(dirname); err != nil {
		return err
	}

	return nil
}

func (r *ResourceController) GetFileInformations(dirname, filename string) (os.FileInfo, error) {
	var (
		fileInfo    os.FileInfo
		err         error
		firstChange changeWorkDir = changeWorkerDirectory
		lastChange  backToHomeDir = changeToMainDirectory
	)

	defer lastChange()

	if err := firstChange(dirname); err != nil {
		return fileInfo, err
	}

	fileInfo, err = os.Stat(filename)
	if err != nil {
		return fileInfo, err
	}
	return fileInfo, nil
}

func (r *ResourceController) WriteRemoteCSV(dir, filename, queryType string, query []string) error {
	var (
		firstChange changeWorkDir = changeWorkerDirectory
		lastChange  backToHomeDir = changeToMainDirectory
	)

	if queryType != insert {
		return errors.New("Invalid Crud Operation")
	}

	defer lastChange()
	if err := firstChange(dir); err != nil {
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

	return nil
}

// TODO
func (r *ResourceController) ReadInRemoteCSV(dir, filename, queryType string, query []string) error {

	return nil
}

func (r *ResourceController) UpdateRemoteCSV(dir, filename, queryType string, query map[string][]string) error {
	var (
		firstChange changeWorkDir = changeWorkerDirectory
		lastChange  backToHomeDir = changeToMainDirectory
	)

	if queryType != update {
		return errors.New("Invalid Crud Operation")
	}

	defer lastChange()
	if err := firstChange(dir); err != nil {
		return err
	}

	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	fileContent, err := reader.ReadAll()
	if err != nil {
		return err
	}
	// parse the file content
	for _, value := range fileContent {
		// TODO
	}
	return nil
}

// TODO
func (r *ResourceController) DeleteRemoteCSV(dir, filename string, query []string) error {
	return nil
}

func changeWorkerDirectory(dirname string) error {
	switch {
	case dirname != "":
		for {
			if err := os.Chdir(filepath.Join("local_file_system", dirname)); err != nil {
				return err
			} else {
				break
			}
		}
	case dirname == "":
		for {
			if err := os.Chdir(filepath.Join("local_file_system")); err != nil {
				return err
			} else {
				break
			}
		}
	}

	return nil
}

func changeToMainDirectory() error {
	if err := os.Chdir("E:/RenzoFS"); err != nil {
		return err
	}
	return nil
}
