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
	IPaddress            string
	ListenPortForClient  string
	ListenPortForServers string
	TCPInfo              string
	// contain servers to ruote
	// request to
	serverPool []ServerPool
	clientConn net.Conn
	serverConn net.Conn
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
func (s *RenzoFSReverseProxy) NewServer(ipAddress, tcpInfo, clientListenPort, serverListenPort string) *RenzoFSReverseProxy {
	return &RenzoFSReverseProxy{
		IPaddress:            ipAddress,
		ListenPortForClient:  clientListenPort,
		ListenPortForServers: serverListenPort,
		TCPInfo:              tcpInfo,
		serverPool:           make([]ServerPool, 0),
		clientConn:           nil,
		serverConn:           nil,
	}
}

// this method make the server listen to
// incoming request from clients
func (s *RenzoFSReverseProxy) StartListenForClient() {
	var err error
	completeAddress := makeEntireReverseProxyServerAddress(s.ListenPortForClient, s.IPaddress)
	listener, err := net.Listen(s.TCPInfo, completeAddress)
	if err != nil {
		log.Fatal(err)
	}

	for {
		s.clientConn, err = listener.Accept()
		if err != nil {
			log.Fatal(err)
			break
		}
		go handleRequests(s.clientConn)
	}
}

func (s *RenzoFSReverseProxy) StartListenForServers() {
	var err error
	completeAddress := makeEntireReverseProxyServerAddress(s.ListenPortForServers, s.IPaddress)
	listener, err := net.Listen(s.TCPInfo, completeAddress)
	if err != nil {
		log.Fatal(err)
	}

	for {
		s.serverConn, err = listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go checkServerPool(s.serverConn)
	}
}

// this method close the connection beetwen
// the reverse proxy server to clients and
// other servers
func (s *RenzoFSReverseProxy) CloseAll() {
	s.clientConn.Close()
	s.serverConn.Close()
}

// this function enable reverse proxy server
// to handle both clients and server request
func handleRequests(req net.Conn) {
	// TODO
}

func checkServerPool(conn net.Conn) {

}

func addToServerPool() {
	// TODO
}

// this function couple the host name to
// the port name
func makeEntireReverseProxyServerAddress(port, ipAddress string) string {
	return net.JoinHostPort(ipAddress, port)
}
