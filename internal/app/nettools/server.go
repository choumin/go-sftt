package nettools

import (
	"net"
	"../common"
	"strings"
	"fmt"
	"strconv"
	"../fileops"
	"../decode"
)

type Server struct {

}

func (s *Server) Init() {
	ip := ""
	port := GeneratePort()
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	fmt.Println(addr)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	common.CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	common.CheckError(err)
	request := make([]byte, fileops.BUF_SIZE)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		_, err = conn.Read(request)
		if err != nil {
			fmt.Println(err)
			continue
		}

		//fmt.Println(string(request))
		decode.Decode(request)
		request = make([]byte, fileops.BUF_SIZE)
	}
}

func NewServer() (s *Server) {
	s = &Server{}
	return 
}