package authservice

import (
	"database/sql"
	"net/http"
)

type DeleteFriendship struct {
	user, friend string
	db           *sql.DB
}

func (d *DeleteFriendship) HandleFriendshipElimination(w http.ResponseWriter, r *http.Request) {

}

func (d *DeleteFriendship) openConnection(w http.ResponseWriter) {

}

func (d *DeleteFriendship) deleteStatement(w http.ResponseWriter) {

}
