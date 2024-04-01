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
		firstChange     changeWorkDir = changeWorkerDirectory
		lastChange      backToHomeDir = changeToMainDirectory
		positionalIndex []int         = make([]int, 0)
		columns         []string      = make([]string, 0)
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
	// da migliorare
	for index, value := range fileContent {
		for index := range value {
			for key := range query {
				switch {
				case value[index] == key:
					positionalIndex = append(positionalIndex, index) // indici che mi servono per muovermi dopo
					columns = append(columns, key)                   // chiavi per prendere i valori della mappa
				}
			}
		}

		if index >= 1 {
			for index := range value {
				for _, position := range positionalIndex {
					if index == position {
						valueSet := query[columns[index]]
						if valueSet[0] == value[index] {
							value[index] = valueSet[1] // modifico il dato
						}
					}
				}
			}
		}
	}

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(fileContent); err != nil {
		return err
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
