/**
*	@author Elia Renzoni
*	@date 20/04/2024
*	@brief simple client
 */

package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":3030")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("uazzap server !")); err != nil {
		log.Fatal(err)
	}
}
