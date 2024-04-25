/**
*	@author Elia Renzoni
*	@date 20/04/2024
*	@brief v1 of RenzoFS reverse proxy server
 */

package main

import (
	renzofsreverseproxy "renzofs/renzofs_reverse_proxy"
)

func main() {
	reverseProxy := renzofsreverseproxy.NewReverseProxyServer("localhost", "tcp", ":5050", ":6060")
	reverseProxy.StartListenForClient()
}
