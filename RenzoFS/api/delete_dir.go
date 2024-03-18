package api

import "net/http"

type DeleteDirPayLoad struct {
  dirToDeleteName string
}

func (d *DeleteDirPayLoad) HandleDirElimination(w http.ResponseWriter, r *http.Request) {

}
