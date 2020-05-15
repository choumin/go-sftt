package main

import (
	"../../nettools"
)
func main() {
	server := nettools.NewServer()
	listener := server.Init()
	server.Run(listener)
}
