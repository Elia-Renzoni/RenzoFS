package main

import (
	"net/http"
	statistichandler "renzofs/statistics_service/statistic_handler"
	"time"
)

func main() {
	stat := &statistichandler.StatPayLoadstruct{}
	health := &statistichandler.HealthSystems{}

	router := http.NewServeMux()
	router.HandleFunc("/statistics/", stat.HandleRead)
	router.HandleFunc("/health", health.HandleHealthCheck)

	statsServer := &http.Server{
		Addr:              ":8081",
		ReadTimeout:       3 * time.Second,
		WriteTimeout:      3 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router,
	}

	statsServer.ListenAndServe()
}
