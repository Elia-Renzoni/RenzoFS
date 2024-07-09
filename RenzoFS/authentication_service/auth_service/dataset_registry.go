package authservice

import (
	"database/sql"
	"net/http"
)

type DataSetRegistry struct {
	folder string `json:"new_folder"`
	user   string `json:"user"`
	db     *sql.DB
}

func (d *DataSetRegistry) HandleRegistry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {

	}

}

func (d *DataSetRegistry) openConnection(w http.ResponseWriter) {
	var err error
	conn := "postgres://postgres:elia@localhost/renzofsdb?sslmode=disable"
	d.db, err = sql.Open("postgres", conn)

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	if err := d.db.Ping(); err != nil {
		http.Error(w, "Ping Error", 500)
	}
}

// TODO
func (d *DataSetRegistry) insertStatement(w http.ResponseWriter) {
	insert := `INSERT INTO folders(folder, user_id)
	           VALUES ($1, $2);`

}
