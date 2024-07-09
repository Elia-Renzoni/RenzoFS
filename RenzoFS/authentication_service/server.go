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

	router.HandleFunc("/signin", signin.HandleSignIn)
	router.HandleFunc("/signout", signout.HandleSignout)
	router.HandleFunc("/registry", registry.HandleRegistry)
	router.HandleFunc("/deregistry/", deregistry.HandleDeregistry)

	http.ListenAndServe(":8082", router)
}
