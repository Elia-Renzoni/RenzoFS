/**
*	@author Elia Renzoni
*	@date 20/04/2024
*	@brief RenzoFS reverse proxy
**/

package renzofsreverseproxy

import (
	"log"
	"net"
)

type RenzoFSReverseProxy struct {
	IPaddress string
	Port      string
	TCPInfo   string
	// contain servers to ruote
	// request to
	serverPool []ServerPool
	conn       net.Conn
}

type ServerPool struct {

	// server api endpoint
	endpoint string

	// server listen port
	port int

	// server ip address
	IPaddress string
}

// this function set-up a new reverse proxy server
// by passing as arguments the ip adress and the
// port by witch the server can start
func (s *RenzoFSReverseProxy) NewServer(ipAddress, tcpInfo, listenPort string) *RenzoFSReverseProxy {
	return &RenzoFSReverseProxy{
		IPaddress:  ipAddress,
		Port:       listenPort,
		TCPInfo:    tcpInfo,
		serverPool: make([]ServerPool, 0),
		conn:       nil,
	}
}

// this method make the server listen to
// incoming request from clients
func (s *RenzoFSReverseProxy) Start() {
	var err error
	completeAddress := makeEntireReverseProxyServerAddress(s.Port, s.IPaddress)
	listener, err := net.Listen(s.TCPInfo, completeAddress)
	if err != nil {
		log.Fatal(err)
	}

	for {
		s.conn, err = listener.Accept()
		if err != nil {
			log.Fatal(err)
			break
		}
		go handleRequests(s.conn)
	}
}

// this method close the connection beetwen
// the reverse proxy server to clients and
// other servers
func (s *RenzoFSReverseProxy) Close() {
	s.conn.Close()
}

// this function enable reverse proxy server
// to handle both clients and server request
func handleRequests(req net.Conn) {
	// TODO
}

func addToServerPool() {
	// TODO
}

// this function couple the host name to
// the port name
func makeEntireReverseProxyServerAddress(port, ipAddress string) string {
	return net.JoinHostPort(ipAddress, port)
}
