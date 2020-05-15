package common

import (
	"fmt"
	"os"
)

const (
	FILE_TYPE_UNKNOW = 0
	FILE_TYPE_FILE = 1
	FILE_TYPE_DIR = 2

	FILE_TYPE_MAGIC_LEN = 2
	FILE_TYPE_MAGIC_FILE uint16 = 0x0930
	FILE_TYPE_MAGIC_DIR uint16 = 0x1991

	MAX_FILE_NAME_LEN = 256
)

func CheckError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func GetFileType(path string) (fileType int, err error) {
	fi, e := os.Stat(path)
	if e != nil {
		return FILE_TYPE_UNKNOW, e
	}
	if fi.IsDir() {
		return FILE_TYPE_DIR, nil
	}

	return FILE_TYPE_FILE, nil
}