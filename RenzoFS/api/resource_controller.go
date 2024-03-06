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
)

type ResourceController struct {
}

/*func (r *ResourceController) checkDir(name string) (result bool) {
	if dirInfo, err := os.Stat(name); dirInfo.IsDir() {
		result = true
	} else if dirInfo != nil {
		fmt.Printf("Qualcosa di strano è successa")
	} else if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("La Directory non esiste")
			result = false
		}
		log.Fatal(err)
	}
	return
}

func (r *ResourceController) checkFile(name string) (result bool) {
	if fileInfo, err := os.Stat(name); fileInfo.Name() != name { // to test
		result = true
	} else if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Ïl File non esiste")
			result = false
		}
	}
	return
}*/

func (r *ResourceController) createNewDir(dirname string) {
	if err := os.MkdirAll("renzofs_local_file_system/"+dirname, os.ModeDir); err != nil {
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
