package main

import (
	"net/http"
	statistichandler "renzofs/statistics_service/statistic_handler"
)

func main() {
	stat := &statistichandler.StatPayLoadstruct{}
	router := http.NewServeMux()
	router.HandleFunc("/statistics", stat.HandleRead)

	http.ListenAndServe(":8081", router)
}
