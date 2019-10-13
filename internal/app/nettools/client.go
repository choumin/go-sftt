package nettools

import (
	"net"
	"../common"
	"strings"
	"strconv"
	"fmt"
)
type Client struct {

}

func (c *Client) Init() {
	ip := "127.0.0.1"
	port := GeneratePort()
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	fmt.Println(addr)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	common.CheckError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	common.CheckError(err)
	_, err = conn.Write([]byte("hello server"))
	common.CheckError(err)
}
func (c *Client) Connect() {

}
func NewClient() (c *Client) {
	c = &Client{}
	return
}
