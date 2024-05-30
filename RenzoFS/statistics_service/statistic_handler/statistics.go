package statistichandler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type StatPayLoadstruct struct {
	dirname, filename string
}

func (s *StatPayLoadstruct) HandleRead(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path
	splittedRequest := strings.Split(requestPath, "/")
	s.dirname = splittedRequest[1]
	s.filename = splittedRequest[2]
	fmt.Printf("%v - %v ", s.dirname, s.filename)

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		// contact the remote storage service
		transport := &http.Transport{
			IdleConnTimeout: 30 * time.Second,
		}
		microservice := http.Client{
			Transport: transport,
		}
		result, err := url.JoinPath("localhot:8080", "/read", s.dirname, s.filename)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		resp, err := microservice.Get(result)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		w.WriteHeader(resp.StatusCode)
		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		io.Copy(w, resp.Body)
	}
}
