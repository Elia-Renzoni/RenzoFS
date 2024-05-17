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
	apiGateway.StartListeningRequests()
}
