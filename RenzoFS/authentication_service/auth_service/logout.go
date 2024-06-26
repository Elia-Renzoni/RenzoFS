package authservice

import "net/http"

type Logout struct {
}

func (l *Logout) HandleLogout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FOO Logout"))
}
