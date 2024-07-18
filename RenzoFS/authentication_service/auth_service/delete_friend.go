package authservice

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

type DeleteFriendship struct {
	user, friend string
	db           *sql.DB
}

func (d *DeleteFriendship) HandleFriendshipElimination(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		splittedRequest := strings.Split(r.URL.Path, "/")
		d.user = splittedRequest[2]
		d.friend = splittedRequest[3]

		d.openConnection(w)
		defer d.db.Close()

		d.deleteStatement(w)
	}
}

func (d *DeleteFriendship) openConnection(w http.ResponseWriter) {
	var err error
	conn := "postgres://elia:elia@localhost/renzofsdb?sslmode=disable"
	d.db, err = sql.Open("postgres", conn)

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	if err := d.db.Ping(); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (d *DeleteFriendship) deleteStatement(w http.ResponseWriter) {
	deleteFriendshipStmt := `DELETE FROM friends
							 WHERE user1 = $1 AND user2 = $2`

	_, err := d.db.Exec(deleteFriendshipStmt, d.user, d.friend)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		jsonMessage, err := json.Marshal(map[string]string{
			"succ_message": "friendship succesfully deleted",
		})
		if err != nil {
			http.Error(w, "Marhaling error", 500)
		} else {
			w.WriteHeader(http.StatusAccepted)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonMessage)
		}
	}
}
