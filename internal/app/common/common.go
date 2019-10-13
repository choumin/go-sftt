package common

import (
	"fmt"
	"os"
)
func CheckError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}