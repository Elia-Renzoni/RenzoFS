package statistichandler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"
	"time"
)

type StatPayLoadstruct struct {
	dirname, filename string
}

func (s *StatPayLoadstruct) HandleRead(w http.ResponseWriter, r *http.Request) {
	_, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	requestPath := r.URL.Path
	splittedRequest := strings.Split(requestPath, "/")
	s.dirname = splittedRequest[2]
	s.filename = splittedRequest[3]
	fmt.Printf("%v - %v ", s.dirname, s.filename)

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		fmt.Printf(path.Join("localhost:8080", "stats", s.dirname, s.filename))
		resp, err := http.Get(path.Join("localhost:8080", "stats", s.dirname, s.filename))
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		w.WriteHeader(resp.StatusCode)
		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		io.Copy(w, resp.Body)
	}
}
