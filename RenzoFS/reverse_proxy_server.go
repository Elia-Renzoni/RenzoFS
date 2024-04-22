/**
*	@author Elia Renzoni
*	@date 20/04/2024
*	@brief v1 of RenzoFS reverse proxy server
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	buffer := make([]byte, 1024)
	listener, err := net.Listen("tcp", ":3030")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		if _, err := conn.Read(buffer); err != nil || err == io.EOF {
			break
		}
		fmt.Printf("%s", string(buffer))
		conn.Close()
	}
}
