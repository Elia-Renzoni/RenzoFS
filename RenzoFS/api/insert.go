/**
*	@author Elia Renzoni
*	@date 02/02/2024
*	@brief Insert json query to files API
**/

package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type InsertPayLoad struct {
	User         string                 `json:"user"`
	FileName     string                 `json:"file_name"`
	QueryContent map[string]interface{} `json:"query_content"`
}

var payload InsertPayLoad = InsertPayLoad{}

func HandleInsertion(w http.ResponseWriter, r *http.Request) {
	var errMessage ResponseMessages = ResponseMessages{}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json, _ := errMessage.MarshallErrMessage()
		w.Write(json)
	} else {
		defer r.Body.Close()
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &payload)
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	}
	printContent()
}

func printContent() {
	fmt.Printf("User : %s\n", payload.User)
	fmt.Printf("FileName: %s\n", payload.FileName)
	for key, value := range payload.QueryContent {
		fmt.Printf("Key : %s", key)
		switch eff := value.(type) {
		case string:
			fmt.Printf("Value : %s", eff)
		case int:
			fmt.Printf("Value : %d", eff)
		case float64:
			fmt.Printf("Value: %f", eff)
		}
	}
}
