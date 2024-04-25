/**
*	@author Elia Renzoni
*	@date 20/04/2024
*	@brief RenzoFS reverse proxy
**/

package renzofsreverseproxy

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
	port string

	// server ip address
	IPaddress string
}

// this function set-up a new reverse proxy server
// by passing as arguments the ip adress and the
// port by witch the server can start
func (s *RenzoFSReverseProxy) NewReverseProxyServer(ipAddress, tcpInfo, clientListenPort, serverListenPort string) *RenzoFSReverseProxy {
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

// this method make the server listen
// to incoming service request that
// wuold add theif informations in the
// server pool Slice
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
		checkServerPool(s.serverConn, s)
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
func handleRequests(conn net.Conn) {
	// TODO
}

// this function must control if the
// service information are in the slice
// if there are not in the slice it must
// add it.
// the message send from the services
// follow this form: enpoint \t addr \t port
func checkServerPool(conn net.Conn, s *RenzoFSReverseProxy) {
	var (
		message string
		scanner *bufio.Scanner
		exist   bool = true
	)
	scanner = bufio.NewScanner(conn)
	// takes one message from services each iteration
	for scanner.Scan() {
		message = scanner.Text()
		splittedMessage := strings.Split(message, "\t")
		endpoint := splittedMessage[0]
		address := splittedMessage[1]
		port := splittedMessage[2]

		for index := range s.serverPool {
			switch {
			case !(s.serverPool[index].endpoint == endpoint):
				fallthrough
			case !(s.serverPool[index].IPaddress == address):
				fallthrough
			case !(s.serverPool[index].port == port):
				exist = false
			default:
				exist = true
			}
		}

		// if informations doesn't exist
		if !exist {
			addToServerPool(endpoint, address, port, s)
		}
	}
}

// this function add the services to
// server pool slice
func addToServerPool(endpoint, address, port string, s *RenzoFSReverseProxy) {
	s.serverPool = append(s.serverPool, ServerPool{
		endpoint:  endpoint,
		IPaddress: address,
		port:      port,
	})
	fmt.Printf("Added a new Service to Server Pool")
}

// this function couple the host name to
// the port name
func makeEntireReverseProxyServerAddress(port, ipAddress string) string {
	return net.JoinHostPort(ipAddress, port)
}
