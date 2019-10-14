package nettools

import (
	"net"
	"../common"
	"strings"
	"strconv"
	"fmt"
	"../encode"
)

type Client struct {
	ip string
	port int
	conn *net.TCPConn
}

func (c *Client) Init() {
	c.ip = "127.0.0.1"
	c.port = GeneratePort()
	
}
func (c *Client) Connect() {
	addr := strings.Join([]string{c.ip, strconv.Itoa(c.port)}, ":")
	fmt.Println(addr)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	common.CheckError(err)
	c.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	common.CheckError(err)
	
}
func (c *Client) Send(path string) {
	buf := encode.Encode(path)
	c.conn.Write(buf[0:len(buf)])
}
func NewClient() (c *Client) {
	c = &Client{}
	return
}
