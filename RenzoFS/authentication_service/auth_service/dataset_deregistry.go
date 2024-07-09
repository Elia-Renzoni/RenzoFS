package authservice

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

type DataSetDeregistry struct {
	username string
	db       *sql.DB
}

func (d *DataSetDeregistry) HandleDeregistry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		d.openConnection(w)
		defer d.db.Close()

		splittedRequest := strings.Split(r.URL.Path, "/")
		d.username = splittedRequest[2]

		d.deleteStatement(w)
	}
}

func (d *DataSetDeregistry) openConnection(w http.ResponseWriter) {
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

func (d *DataSetDeregistry) deleteStatement(w http.ResponseWriter) {
	delete := `DELETE FROM folders
			   WHERE user_id = $1`
	_, err := d.db.Exec(delete, d.username)

	if err != nil {
		http.Error(w, "Statement Error", 500)
	} else {
		jsonMessage, err := json.Marshal(map[string]string{
			"succ_message": "Folders Succesfully Deleted",
		})
		if err != nil {
			http.Error(w, "Marshaling Error", 500)
		} else {
			w.WriteHeader(http.StatusAccepted)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonMessage)
		}
	}
}
