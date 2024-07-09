/**
*	@author Elia Renzoni
*	@date 20/04/2024
*	@brief v1 of RenzoFS reverse proxy server
 */

package main

import (
	renzofsapigateway "renzofs/renzofs_api_gateway"
)

func main() {
	apiGateway := renzofsapigateway.NewRenzoFSAPIGateway("localhost", ":4040")
	apiGateway.AddMicroservice("read", "http://127.0.0.1:8080",
		"insert", "http://127.0.0.1:8080",
		"update", "http://127.0.0.1:8080",
		"delete", "http://127.0.0.1:8080",
		"deletedir", "http://127.0.0.1:8080",
		"createdir", "http://127.0.0.1:8080",
		"fileinfo", "http://127.0.0.1:8080",
		"statistics", "http://127.0.0.1:8081",
		"login", "http://127.0.0.1:8082",
		"logout", "http://127.0.0.1:8082",
		"registry", "http://127.0.0.1:8082",
		"deregistry", "http://127.0.0.1:8082")
	apiGateway.StartListeningRequests()
}
