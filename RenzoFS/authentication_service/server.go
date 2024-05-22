package main

import "net/http"

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/login")
	router.HandleFunc("/logout")

	http.ListenAndServe(":8082", router)
}
