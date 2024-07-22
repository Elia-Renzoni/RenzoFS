package renzofsapigateway

import (
	"net/http"
	"sync"
	"time"
)

type ServiceDiscovery struct {
	services []string
	RenzoFSAPIGateway
	mutex sync.Mutex
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
	for _, value := range s.services {
		go func(server string) {
			var decodedRes string
			client := &http.Client{
				Timeout: 3 * time.Second,
			}
			response, err := client.Get(server)
			// TODO
		}(value)
	}
}
