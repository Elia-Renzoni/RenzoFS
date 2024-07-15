package authservice

import (
	"database/sql"
	"encoding/json"
	"io"
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
		d.openConnection(w)
		defer d.db.Close()

		reqBody, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		json.Unmarshal(reqBody, d)

		d.insertStatement(w)
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

func (d *DataSetRegistry) insertStatement(w http.ResponseWriter) {
	insertion := `INSERT INTO folders(folder_name, username)
	        		VALUES ($1, $2);`
	_, err := d.db.Exec(insertion, d.folder, d.user)
	if err != nil {
		http.Error(w, "Statement Error", 500)
	} else {
		jsonMessage, err := json.Marshal(map[string]string{
			"succ_message": "folder succesfully added",
		})
		if err != nil {
			http.Error(w, "Marshaling Error", 500)
		} else {
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonMessage)
		}
	}
}
