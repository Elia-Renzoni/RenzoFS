package api

import (
	"encoding/json"
	"net/http"
)

type HealthSystems struct {
}

func (h *HealthSystems) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		json, _ := json.Marshal(map[string]string{"port_name": "Method Not Allowed"})
		http.Error(w, string(json), http.StatusMethodNotAllowed)
	} else {
		jsonMessage, err := json.Marshal(map[string]string{
			"port_name": "8080",
		})
		if err != nil {
			json, _ := json.Marshal(map[string]string{
				"port_name": "" + err.Error(),
			})
			http.Error(w, string(json), 500)
		} else {
			w.WriteHeader(http.StatusAccepted)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonMessage)
		}
	}
}
