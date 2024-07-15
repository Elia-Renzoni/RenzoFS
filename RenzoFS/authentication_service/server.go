package main

import (
	"net/http"
	renzofs "renzofs/authentication_service/auth_service"
)

func main() {
	router := http.NewServeMux()
	signin := &renzofs.SignIn{}
	signout := &renzofs.Signout{}
	registry := &renzofs.DataSetRegistry{}
	deregistry := &renzofs.DataSetDeregistry{}
	newFriend := &renzofs.NewFriendship{}
	deleteFriend := &renzofs.DeleteFriendship{}

	router.HandleFunc("/signin", signin.HandleSignIn)
	router.HandleFunc("/signout", signout.HandleSignout)
	router.HandleFunc("/registry", registry.HandleRegistry)
	router.HandleFunc("/deregistry/", deregistry.HandleDeregistry)
	router.HandleFunc("/newfriend", newFriend.HandleNewFriendship)
	router.HandleFunc("/deletefriend", deleteFriend.HandleFriendshipElimination)

	http.ListenAndServe(":8082", router)
}
