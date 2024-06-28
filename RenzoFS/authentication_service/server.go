package main

import (
	"net/http"
	renzofs "renzofs/authentication_service/auth_service"
)

func main() {
	router := http.NewServeMux()
	login := &renzofs.Login{}
	logout := &renzofs.Logout{}
	registry := &renzofs.DataSetRegistry{}
	deregistry := &renzofs.DataSetDeregistry{}
	addFriend := &renzofs.AddFriend{}
	deleteFriend := &renzofs.DeleteFriend{}

	router.HandleFunc("/login", login.HandleLogin)
	router.HandleFunc("/logout", logout.HandleLogout)
	router.HandleFunc("/registry", registry.HandleRegistry)
	router.HandleFunc("/deregistry/", deregistry.HandleDeregistry)
	router.HandleFunc("/newf", addFriend.HandleFriendAdding)
	router.HandleFunc("/deletef", deleteFriend.HandleFriendElimination)

	http.ListenAndServe(":8082", router)
}
