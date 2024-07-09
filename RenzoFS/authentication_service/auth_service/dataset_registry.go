package authservice

import (
	"database/sql"
	"net/http"
)

type DataSetRegistry struct {
	folder string `json:"new_folder"`
	db     *sql.DB
}

func (d *DataSetRegistry) HandleRegistry(w http.ResponseWriter, r *http.Request) {

}

func (d *DataSetRegistry) openConnection(w http.ResponseWriter) {

}

func (d *DataSetRegistry) insertStatement(w http.ResponseWriter) {

}
