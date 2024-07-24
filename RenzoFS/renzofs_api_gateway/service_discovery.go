package renzofsapigateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	port1, port2, port3 string = "8080", "8081", "8082"
)

type ServiceDiscovery struct {
	services []string
	RenzoFSAPIGateway
}

type HealthCheckRes struct {
	PortName string `json:"port_name"`
}

func NewServiceDiscovery() ServiceDiscovery {
	return ServiceDiscovery{}
}

func (s *ServiceDiscovery) takeValuesAndTrim() {
	var (
		i    int
		oldV string
	)

	for _, microservice := range s.serverPool {
		if i != 0 {
			if oldV != microservice {
				s.services = append(s.services, microservice)
			}
		}
		oldV = microservice
		i++
	}
}

func (s *ServiceDiscovery) Broadcast() {
	for _, microservice := range s.services {
		go s.findMicroservices(microservice)
	}
}

func (s *ServiceDiscovery) findMicroservices(server string) {
	urlReq := fmt.Sprintf("%v/%v", server, "health")

	for {
		time.Sleep(5 * time.Second)

		client := &http.Client{
			Timeout: 3 * time.Second,
		}
		response, err := client.Get(urlReq)

		if err != nil {
			if os.IsTimeout(err) {
				s.deleteMicroservice(server)
			}
		} else {
			s.addMicroservice(response, server)
		}
	}
}

func (s *ServiceDiscovery) deleteMicroservice(server string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for endpoint, microservice := range s.serverPool {
		if microservice == server {
			parsed, _ := url.Parse(microservice)
			s.serverPool[endpoint] = string(parsed.Port())
		}
	}
}

func (s *ServiceDiscovery) addMicroservice(response *http.Response, server string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	encoded := &HealthCheckRes{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, encoded)

	if response.StatusCode != http.StatusServiceUnavailable {
		for endpoint, microservice := range s.serverPool {
			switch microservice {
			case port1, port2, port3:
				s.serverPool[endpoint] = server
			}
		}
	} else {
		s.deleteMicroservice(server)
	}
}
