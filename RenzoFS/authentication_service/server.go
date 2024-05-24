package main

import (
	"net/http"
	renzofs "renzofs/authentication_service/auth_service"
)

func main() {
	router := http.NewServeMux()
	auth := &renzofs.Login{}
	auth1 := &renzofs.Logout{}

	router.HandleFunc("/login", auth.HandleLogin)
	router.HandleFunc("/logout", auth1.HandleLogout)

	http.ListenAndServe(":8082", router)
}
