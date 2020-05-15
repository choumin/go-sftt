package fileops

import (
	"fmt"
	"testing"
)

func TestGenerateMetaFileContent(t *testing.T) {
	path := "/User/didi/Documents/dev/tmp.txt"
	metaFileContent := GenerateMetaFileContent(path)
	fmt.Println(len(metaFileContent))
	//for _, byte := range metaFileContent {
	//	fmt.Print(byte)
	//}
	//tmp := string(metaFileContent)
	//fmt.Println(tmp + "xx")
}
