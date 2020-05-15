package nettools

import (
	"net"
	"../common"
	"strings"
	"fmt"
	"strconv"
	"../fileops"
	"../decode"
	"../store"
)

type Server struct {

}

func (s *Server) Init() (*net.TCPListener) {
	ip := ""
	port := GeneratePort()
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	fmt.Println(addr)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	common.CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	common.CheckError(err)

	return listener
}

func (s *Server) isSessionBegin(bytes []byte) (bool) {

	return false
}
func (s *Server) isSessionEnd(bytes []byte) (bool) {

	return false
}
func (s *Server) handleRequest(conn net.Conn) {
	defer conn.Close()
	dir := store.CreateDir()
	request := make([]byte, fileops.BUF_SIZE)
	for {
		n, err := conn.Read(request[0:])
		if err != nil {
			break
		}
		decode.Decode(request[0 : n])
		if s.isSessionEnd(request[0 : n]) {
			break
		}
		store.SaveToFile(request[0 : n])
	}
	store.Restore(dir)
}
func (s *Server) Run(listener *net.TCPListener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go s.handleRequest(conn)

	}
}

func NewServer() (s *Server) {
	s = &Server{}
	return 
}