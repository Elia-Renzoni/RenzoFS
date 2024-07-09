package renzofsapigateway

// il servizio deve periodicamente mandare una richiesta
// ai server che si trovano nel server pool.
// Legge il server pool e prende la porta e l'host,
// poi invia a quell'host una richeista all'endpoint
// health. Il server in questione deve rispondere
// con lo stato del sistema, In risposta i server
// devono mandare lo stato della connessione
// come risposta. Se i server non inviano in tempo
// lo stato (da notare ogni richiesta ha un tempo di
// massimo 5 secondi), allora singifica che il server è in down
// e quindi va eliminato dal server pool.

type ServiceDiscovery struct {
	RenzoFSAPIGateway
	toDelete <-chan string
}

func (s *ServiceDiscovery) ManageServerPool() {

}

func visitServerPool(pool map[string]string) bool {
	return false
}

func deleteServerFromPool(keys ...string) {

}
