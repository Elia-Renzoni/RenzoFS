package authservice

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
)

type SignIn struct {
	username string `json:"username"`
	password string `json:"password"`
	db       *sql.DB
}

func (s *SignIn) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		s.openConnection(w)
		defer s.db.Close()

		reqBody, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		json.Unmarshal(reqBody, s)

		s.insertStatement(w)
	}
}

func (s *SignIn) openConnection(w http.ResponseWriter) {
	var err error
	conn := "postgres://postgres:elia@localhost/renzofsdb?sslmode=disable"
	s.db, err = sql.Open("postgres", conn)

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
	}

	if err := s.db.Ping(); err != nil {
		http.Error(w, "Ping Error", 500)
	}
}

func (s *SignIn) insertStatement(w http.ResponseWriter) {
	insert := `INSERT INTO users(username, password)
	           VALUES ($1, $2);`
	_, err := s.db.Exec(insert, s.username, s.password)

	if err != nil {
		http.Error(w, "Statement Error", 500)
	} else {
		succResponse := map[string]string{
			"succ_message": "Account Succesfully Created",
		}
		jsonMessage, err := json.Marshal(succResponse)
		if err != nil {
			http.Error(w, "Marshaling Error", 500)
		} else {
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonMessage)
		}
	}
}
