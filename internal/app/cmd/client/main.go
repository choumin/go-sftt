package main

import (
	"../../nettools"
	"../../common"
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
	fileType, err := common.GetFileType(path)
	if err != nil {
		fmt.Print("get file type failed!")
	}
	client := nettools.NewClient()
	client.Init()
	client.Connect()
	//client.Send(path)
	switch fileType {
	case common.FILE_TYPE_FILE:
		client.SendFile(path)
		break
	case common.FILE_TYPE_DIR:
		client.SendDir(path)
		break
	}
}
