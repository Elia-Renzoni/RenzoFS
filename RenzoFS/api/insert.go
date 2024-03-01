/**
*	@author Elia Renzoni
*	@date 01/03/2024
*
*	@brief Insert in file class
*
**/

package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// things to marshall...
const (
	errorMessageTitle string = "error_message"
	errorMessageBody  string = "method not allowed, this operation need a post call"
	messageTitle      string = "message"
	messageBody       string = "added new content to file"
)

type Insert struct {
	User         string                 `json:"user"`
	FileName     string                 `json:"file_name"`
	QueryContent map[string]interface{} `json:"query_content"`
	ErrResponse  map[string]string
	OkResponse   map[string]string
}

var store Insert = Insert{}

// handle the Request, it control if the method request is equal
// to POST. If is not equal to POST it response with a json written
// message error, despite if equal it call the function to insert
// new data in the csv file
func (i *Insert) HandleInsertion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		i.ErrResponse = make(map[string]string)
		i.ErrResponse[errorMessageTitle] = errorMessageBody
		jsonResponse, err := json.Marshal(i.ErrResponse)
		if err != nil {
			log.Fatal("error ! ")
		}
		w.Write(jsonResponse)
	} else {
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err := json.Unmarshal(body, &store); err != nil {
			log.Fatal("err")
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		i.OkResponse = make(map[string]string)
		i.OkResponse[messageTitle] = messageBody + i.FileName
		jsonResponse, _ := json.Marshal(i.OkResponse)
		w.Write(jsonResponse)

		// parsing the query content
		insertValues(store.QueryContent)
	}
}

func insertValues(content map[string]interface{}) {
	// cercare se l'utente esiste già o meno
	// se non esite creare la nuova cartella con il file vuoto
	// se già esiste:
	// apro il file il scrittura e scrivo i contenuti

	for key, value := range content {
		// TODO
	}
}
