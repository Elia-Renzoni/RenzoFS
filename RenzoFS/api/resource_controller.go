/**
*	@author Elia Renzoni
*	@date 02/03/2024
*	@brief Check if the directory and files exist
*
**/

package api

import (
	"log"
	"os"
	"path/filepath"
)

type ResourceController struct {
}

func (r *ResourceController) createNewDir(dirname string) {
	if err := os.Mkdir(filepath.Join("local_file_system", dirname), os.ModeDir); err != nil {
		log.Fatal(err)
	}
}

func (r *ResourceController) createNewFile(filename string) (*os.File, error) {
	if file, err := os.OpenFile(filename, os.O_CREATE, 0644); err != nil {
		log.Fatal(err)
	} else {
		return file, err
	}
	return nil, nil
}
