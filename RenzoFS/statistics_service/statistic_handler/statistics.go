package statistichandler

import (
	"io"
	"net/http"
	"path"
	"strings"
)

type StatPayLoadstruct struct {
	dirname, filename string
}

func (s *StatPayLoadstruct) HandleRead(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path
	splittedRequest := strings.Split(requestPath, "/")
	s.dirname = splittedRequest[2]
	s.filename = splittedRequest[3]

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		resp, err := http.Get(path.Join("localhost:8080", "stats", s.dirname, s.filename))
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		w.WriteHeader(resp.StatusCode)
		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		io.Copy(w, resp.Body)
	}
}
