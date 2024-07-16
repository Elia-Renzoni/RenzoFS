package authservice

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	_ "github.com/lib/pq"
)

type NewFriendship struct {
	baseUser string `json:"first_user"`
	friend   string `json:"second_user"`
	db       *sql.DB
}

func (n *NewFriendship) HandleNewFriendship(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		n.openConnection(w)
		defer n.db.Close()

		reqBody, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		json.Unmarshal(reqBody, n)

		n.insertStatement(w)
	}
}

func (n *NewFriendship) openConnection(w http.ResponseWriter) {
	var err error
	conn := "postgres://elia:elia@localhost/renzofsdb?sslmode=disable"
	n.db, err = sql.Open("postgres", conn)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	if err := n.db.Ping(); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (n *NewFriendship) insertStatement(w http.ResponseWriter) {
	insert := `INSERT INTO friends(user1, user2)
				VALUES ($1, $2);`
	_, err := n.db.Exec(insert, n.baseUser, n.friend)
	if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		jsonMessage, err := json.Marshal(map[string]string{
			"succ_message": "Added new friend succesfully",
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
