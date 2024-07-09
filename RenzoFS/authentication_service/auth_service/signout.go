package authservice

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

type Signout struct {
	username string
	db       *sql.DB
}

func (so *Signout) HandleSignout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		so.openConnection(w)
		defer so.db.Close()

		splittedReq := strings.Split(r.URL.Path, "/")
		so.username = splittedReq[2]

		so.deleteStatement(w)
	}
}

func (so *Signout) deleteStatement(w http.ResponseWriter) {
	deleteStatement := `DELETE FROM users
	                    WHERE username = $1;`
	_, err := so.db.Exec(deleteStatement, so.username)

	if err != nil {
		http.Error(w, "Statement Error", 500)
	} else {
		jsonMessage, err := json.Marshal(map[string]string{
			"succ_message": "Account Succesfully Deleted",
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

func (so *Signout) openConnection(w http.ResponseWriter) {
	var err error
	conn := "postgres://postgres:elia@localhost/renzofsdb?sslmode=disable"
	so.db, err = sql.Open("postgres", conn)

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	if err := so.db.Ping(); err != nil {
		http.Error(w, "Ping Error", 500)
	}
}
