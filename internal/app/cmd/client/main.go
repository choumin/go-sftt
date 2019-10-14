package main

import (
	"../../nettools"
	"fmt"
	"os"
)

func help(program string) {
	fmt.Printf("Usage: %s <file | dir>\n", program)
}
func main() {
	if len(os.Args) != 2 {
		help(os.Args[0])
		return
	}
	path := os.Args[1]
	client := nettools.NewClient()
	client.Init()
	client.Connect()
	client.Send(path)
	
	//tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:")
}
