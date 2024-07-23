/**
*	@author Elia Renzoni
*	@date 20/04/2024
*	@brief RenzoFS reverse proxy
**/

package renzofsapigateway

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type RenzoFSAPIGateway struct {
	host       string
	listenPort string
	serverPool map[string]string
	mutex      sync.Mutex
}

// RenzoFS API endpoint
const (
	REMOTE_STORAGE_S_INSERT    string = "insert"
	REMOTE_STORAGE_S_READ      string = "read"
	REMOTE_STORAGE_S_DELETE    string = "delete"
	REMOTE_STORAGE_S_UPDATE    string = "update"
	REMOTE_STORAGE_S_DELETEDIR string = "deletedir"
	REMOTE_STORAGE_S_CREATEDIR string = "createdir"
	REMOTE_STORAGE_S_FILEINFO  string = "fileinfo"

	STATSTIC_S_STATISCTIS string = "statistics"

	AUTHENTICATION_S_LOGIN  string = "login"
	AUTHENTICATION_S_LOGOUT string = "logout"
)

func NewRenzoFSAPIGateway(host, port string) *RenzoFSAPIGateway {
	return &RenzoFSAPIGateway{
		host:       host,
		listenPort: port,
		serverPool: make(map[string]string),
	}
}

func (r *RenzoFSAPIGateway) AddMicroservice(microserviceInfo ...string) {
	for index, mInfo := range microserviceInfo {
		if index >= 1 {
			// index - 1 == key
			// mInfo == value
			r.serverPool[microserviceInfo[index-1]] = mInfo
		}
	}
}

func (r *RenzoFSAPIGateway) StartListeningRequests() {
	router := http.NewServeMux()
	router.HandleFunc("/", r.apiGatewayInnerHandler)

	http.ListenAndServe(r.listenPort, router)
}

func (re *RenzoFSAPIGateway) apiGatewayInnerHandler(w http.ResponseWriter, r *http.Request) {
	var (
		service      *url.URL
		splittedUrl  []string = strings.Split(r.URL.Path, "/")
		endpoint     string   = splittedUrl[1]
		microservice string
		err          error = nil
	)

	// control endpoint
	switch endpoint {
	case REMOTE_STORAGE_S_INSERT:
		if microservice, err = checkServerPool(re.serverPool, REMOTE_STORAGE_S_INSERT); err != nil {
			handleAPIGatewayNegativeResponse(w, err)
		} else {
			service = parseURL(microservice)
		}
	case REMOTE_STORAGE_S_READ:
		if microservice, err = checkServerPool(re.serverPool, REMOTE_STORAGE_S_READ); err != nil {
			handleAPIGatewayNegativeResponse(w, err)
		} else {
			service = parseURL(microservice)
		}
	case REMOTE_STORAGE_S_DELETE:
		if microservice, err = checkServerPool(re.serverPool, REMOTE_STORAGE_S_DELETE); err != nil {
			handleAPIGatewayNegativeResponse(w, err)
		} else {
			service = parseURL(microservice)
		}
	case REMOTE_STORAGE_S_UPDATE:
		if microservice, err = checkServerPool(re.serverPool, REMOTE_STORAGE_S_UPDATE); err != nil {
			handleAPIGatewayNegativeResponse(w, err)
		} else {
			service = parseURL(microservice)
		}
	case REMOTE_STORAGE_S_CREATEDIR:
		if microservice, err = checkServerPool(re.serverPool, REMOTE_STORAGE_S_CREATEDIR); err != nil {
			handleAPIGatewayNegativeResponse(w, err)
		} else {
			service = parseURL(microservice)
		}
	case REMOTE_STORAGE_S_DELETEDIR:
		if microservice, err = checkServerPool(re.serverPool, REMOTE_STORAGE_S_DELETEDIR); err != nil {
			handleAPIGatewayNegativeResponse(w, err)
		} else {
			service = parseURL(microservice)
		}
	case REMOTE_STORAGE_S_FILEINFO:
		if microservice, err = checkServerPool(re.serverPool, REMOTE_STORAGE_S_FILEINFO); err != nil {
			handleAPIGatewayNegativeResponse(w, err)
		} else {
			service = parseURL(microservice)
		}
	case STATSTIC_S_STATISCTIS:
		if microservice, err = checkServerPool(re.serverPool, STATSTIC_S_STATISCTIS); err != nil {
			handleAPIGatewayNegativeResponse(w, err)
		} else {
			service = parseURL(microservice)
		}
	case AUTHENTICATION_S_LOGIN:
		if microservice, err = checkServerPool(re.serverPool, AUTHENTICATION_S_LOGIN); err != nil {
			handleAPIGatewayNegativeResponse(w, err)
		} else {
			service = parseURL(microservice)
		}
	case AUTHENTICATION_S_LOGOUT:
		if microservice, err = checkServerPool(re.serverPool, AUTHENTICATION_S_LOGOUT); err != nil {
			handleAPIGatewayNegativeResponse(w, err)
		} else {
			service = parseURL(microservice)
		}
	default:
		fmt.Printf("******")
		handleAPIGatewayNegativeResponse(w, errors.New("Invalid Endpoint"))
	}

	// spwapping
	r.Host = service.Host
	r.URL.Host = service.Host
	r.URL.Scheme = service.Scheme
	r.RequestURI = ""

	targetServiceResponse, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
		handleAPIGatewayNegativeResponse(w, err)
	} else {
		w.WriteHeader(targetServiceResponse.StatusCode)
		w.Header().Set("Content-Type", targetServiceResponse.Header.Get("Content-Type"))
		io.Copy(w, targetServiceResponse.Body)
	}
}

func handleAPIGatewayNegativeResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadGateway)
	w.Header().Set("Content-Type", "application/json")
	jsonMessage, _ := json.Marshal(map[string]string{"error": err.Error()})
	w.Write(jsonMessage)
}

func checkServerPool(serverPool map[string]string, key string) (string, error) {
	if microservice, ok := serverPool[key]; ok {
		return microservice, nil
	}
	return "", errors.New("Value Not Found")
}

func parseURL(server string) *url.URL {
	target, err := url.Parse(server)
	if err != nil {
		log.Fatal(err)
	}
	return target
}
