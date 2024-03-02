/**
*	@author Elia Renzoni
*	@date 02/03/2024
*	@brief Check if the directory and files exist
*
**/

package api

import "os"

type ResourceController struct {
}

func (r *ResourceController) checkDir(name string) (result bool) {
	_, err := os.Stat(name)
	if controlValue := os.IsExist(err); controlValue {
		result = true
	}
	return
}

func (r *ResourceController) checkFile(name string) (result bool) {
	return
}

func (r *ResourceController) createNewDir(dirname string) {

}

func (r *ResourceController) createNewFile(filename string) {

}
