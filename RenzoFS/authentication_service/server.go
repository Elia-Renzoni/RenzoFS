package main

import (
	"net/http"
	renzofs "renzofs/authentication_service/auth_service"
	"time"
)

func main() {
	router := http.NewServeMux()
	signin := &renzofs.SignIn{}
	signout := &renzofs.Signout{}
	registry := &renzofs.DataSetRegistry{}
	deregistry := &renzofs.DataSetDeregistry{}
	newFriend := &renzofs.NewFriendship{}
	deleteFriend := &renzofs.DeleteFriendship{}
	health := &renzofs.HealthSystems{}

	router.HandleFunc("/signin", signin.HandleSignIn)
	router.HandleFunc("/signout/", signout.HandleSignout)
	router.HandleFunc("/registry", registry.HandleRegistry)
	router.HandleFunc("/deregistry/", deregistry.HandleDeregistry)
	router.HandleFunc("/newfriend", newFriend.HandleNewFriendship)
	router.HandleFunc("/deletefriend/", deleteFriend.HandleFriendshipElimination)
	router.HandleFunc("/health", health.HandleHealthCheck)

	authServer := &http.Server{
		Addr:              ":8082",
		ReadTimeout:       3 * time.Second,
		WriteTimeout:      3 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router,
	}

	authServer.ListenAndServe()
}
