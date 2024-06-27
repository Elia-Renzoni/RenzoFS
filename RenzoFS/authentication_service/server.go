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

	router.HandleFunc("/login", login.HandleLogin)
	router.HandleFunc("/logout", logout.HandleLogout)
	router.HandleFunc("/registry", registry.HandleRegistry)
	router.HandleFunc("/deregistry/", deregistry.HandleDeregistry)

	http.ListenAndServe(":8082", router)
}
