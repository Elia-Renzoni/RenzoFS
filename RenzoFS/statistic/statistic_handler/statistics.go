package statistichandler

import "net/http"

type StatPayLoadstruct struct {
	dirname, filename string
}

func (s *StatPayLoadstruct) HandleRead(w http.ResponseWriter, r *http.Request) {
	// read the request
	// update dirname e filename with the effective path specified by the user
	// make the request
	// wait the response
	// if the response is legit
	//	write it in a json format
	//	send back to the client
	// else
	//	write an error message
}
