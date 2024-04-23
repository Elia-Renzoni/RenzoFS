package statistichandler

import "net/http"

type StatPayLoadstruct struct{}

func (s *StatPayLoadstruct) HandleRead(w http.ResponseWriter, r *http.Request) {
	// TODO
}
