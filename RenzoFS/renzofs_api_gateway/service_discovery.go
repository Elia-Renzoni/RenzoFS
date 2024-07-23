package renzofsapigateway

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
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

	for _, value := range s.serverPool {
		if i != 0 {
			if oldV != value {
				s.services = append(s.services, value)
			}
		}
		oldV = value
		i++
	}
}

func (s *ServiceDiscovery) Broadcast() {
	for _, microservice := range s.services {
		go func(server string) {
			encoded := &HealthCheckRes{}
			client := &http.Client{
				Timeout: 3 * time.Second,
			}
			response, err := client.Get(server)
			if err != nil {
				if os.IsTimeout(err) {
					// devo cancellare subito il servizio
					defer s.mutex.Unlock()
					s.mutex.Lock()

					for endpoint, microservice := range s.serverPool {
						if microservice == server {
							parsed, _ := url.Parse(microservice)
							s.serverPool[endpoint] = string(parsed.Port())
						}
					}
				}
			} else {
				defer s.mutex.Unlock()
				s.mutex.Lock()

				body, _ := io.ReadAll(response.Body)
				json.Unmarshal(body, encoded)

				for endpoint, microservice := range s.serverPool {
					parsed, _ := url.Parse(microservice)
					if parsed.Port() != encoded.PortName {
						if microservice == encoded.PortName {
							s.serverPool[endpoint] = server
						}
					} else {
						// TODO

					}
				}

			}
		}(microservice)
	}
}
