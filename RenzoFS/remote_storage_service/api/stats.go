package api

import (
	"net/http"
	"strings"
)

type StatsPayload struct {
	dirName, fileName string
	messages          *ResponseMessages
	renzoLogger       *RenzoFSCustomLogger
}

func (s *StatsPayload) HandleStats(w http.ResponseWriter, r *http.Request) {
	s.messages = getInstance()
	s.renzoLogger = newRenzoFSCustomLogger()
	parameters := strings.Split(r.URL.Path, "/")
	s.dirName = parameters[2]
	s.fileName = parameters[3]

	if r.Method != http.MethodGet {
		json, err := s.messages.MarshalErrMessage("Method Not Allowed")
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			handleStatsResponses(w, methodNotAllowed, json)
		}
	} else {
		response, err := s.renzoLogger.SearchInLogFile(s.dirName, s.fileName)
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			json, err := s.messages.MarshalSuccesStatResults(response)
			if err != nil {
				http.Error(w, err.Error(), 500)
			} else {
				handleStatsResponses(w, clientSucces, json)
			}
		}
	}
}

func handleStatsResponses(w http.ResponseWriter, id byte, jsonMessage []byte) {
	switch id {
	case serverError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMessage)
	case methodNotAllowed:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMessage)
	case clientSucces:
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonMessage)
	}
}
