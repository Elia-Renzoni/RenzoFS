/**
*	@author Elia Renzoni
*	@date 20/04/2024
*	@brief RenzoFS reverse proxy
**/

package renzofsapigateway

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type RenzoFSAPIGateway struct {
	host       string
	listenPort string
	serverPool map[string]string
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
		err          error
	)

	// control endpoint
	switch endpoint {
	case REMOTE_STORAGE_S_INSERT:
		if microservice, err = checkServerPool(re.serverPool, REMOTE_STORAGE_S_INSERT); err != nil {
			service = parseURL(microservice)
		} else {
			handleAPIGatewayNegativeResponse(w, err)
		}
	case REMOTE_STORAGE_S_READ:
		if microservice, err = checkServerPool(re.serverPool, REMOTE_STORAGE_S_READ); err != nil {
			service = parseURL(microservice)
		} else {
			handleAPIGatewayNegativeResponse(w, err)
		}
	case REMOTE_STORAGE_S_DELETE:
		if microservice, err = checkServerPool(re.serverPool, REMOTE_STORAGE_S_DELETE); err != nil {
			service = parseURL(microservice)
		} else {
			handleAPIGatewayNegativeResponse(w, err)
		}
	case REMOTE_STORAGE_S_UPDATE:
		if microservice, err = checkServerPool(re.serverPool, REMOTE_STORAGE_S_UPDATE); err != nil {
			service = parseURL(microservice)
		} else {
			handleAPIGatewayNegativeResponse(w, err)
		}
	case REMOTE_STORAGE_S_CREATEDIR:
		if microservice, err = checkServerPool(re.serverPool, REMOTE_STORAGE_S_CREATEDIR); err != nil {
			service = parseURL(microservice)
		} else {
			handleAPIGatewayNegativeResponse(w, err)
		}
	case REMOTE_STORAGE_S_DELETEDIR:
		if microservice, err = checkServerPool(re.serverPool, REMOTE_STORAGE_S_DELETEDIR); err != nil {
			service = parseURL(microservice)
		} else {
			handleAPIGatewayNegativeResponse(w, err)
		}
	case REMOTE_STORAGE_S_FILEINFO:
		if microservice, err = checkServerPool(re.serverPool, REMOTE_STORAGE_S_FILEINFO); err != nil {
			service = parseURL(microservice)
		} else {
			handleAPIGatewayNegativeResponse(w, err)
		}
	case STATSTIC_S_STATISCTIS:
		if microservice, err = checkServerPool(re.serverPool, STATSTIC_S_STATISCTIS); err != nil {
			service = parseURL(microservice)
		} else {
			handleAPIGatewayNegativeResponse(w, err)
		}
	default:
		handleAPIGatewayNegativeResponse(w, err)
	}

	// spwapping
	if err != nil {
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
}

func handleAPIGatewayNegativeResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadGateway)
	w.Header().Set("Content-Type", "application/json")
	jsonMessage, _ := json.Marshal(map[string]string{"error": err.Error()})
	w.Write(jsonMessage)
}

func checkServerPool(serverPool map[string]string, key string) (microservice string, err error) {
	var ok bool
	if microservice, ok = serverPool[key]; ok {
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
