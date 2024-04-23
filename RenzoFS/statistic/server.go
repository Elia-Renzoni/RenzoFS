package main

import (
	"net/http"
	statistichandler "renzofs/statistic/statistic_handler"
)

func main() {
	stat := &statistichandler.StatPayLoadstruct{}
	router := http.NewServeMux()
	router.HandleFunc("/statistics", stat.HandleRead)

	http.ListenAndServe(":4040", router)
}
