package fileops

import (
    "os"
    "../common"
    "strings"
    "bufio"
    "fmt"
    "io/ioutil"
)

const BUF_SIZE = 1024
const FILE_NAME_WRAPPER_BEGIN = "FILE_NAME_BEGIN"
const FILE_NAME_WRAPPER_END = "FILE_NAME_END"
const DIR_WRAPPER_FLAG = "DIR"
const TYPE_FILE = 1
const TYPE_DIR = 2 

func IsFile(f string) bool {
	fi, e := os.Stat(f)
    if e != nil {
        return false
    }
    return !fi.IsDir()
}

func generateFileHeader(path string) (header string) {
    return strings.Join([]string{FILE_NAME_WRAPPER_BEGIN, path, FILE_NAME_WRAPPER_END}, "_")
}

func generateDirHeader(path string) (header string) {
    return 
}
func ParseFileType(data []byte) (int) {
    cnt := len(FILE_NAME_WRAPPER_BEGIN)
    if (string(data[0:cnt]) == FILE_NAME_WRAPPER_BEGIN) {
        return TYPE_FILE
    }
    return TYPE_DIR
}
func EncodeFileContent(path string) (data []byte) {
    header := generateFileHeader(path)
    fmt.Printf(header + "\n")

    content, err := ioutil.ReadFile(path)
    common.CheckError(err)
    data = make([]byte, len(header) + len(content) + 1)
    copy(data, []byte(header))
    copy(data[len(header):], content)

    return data
}
func EncodeDirContent(path string) (data []byte) {
    data = make([]byte, 0, BUF_SIZE)
    return
}
func getFilePath(fileName string) (path string) {
    path = fileName
    return
}
func getSfttHomeDir() (homeDir string) {
    return "/Users/choumin/"
}
func DecodeFileContent(data []byte) {
    begin := len(FILE_NAME_WRAPPER_BEGIN) + 1
    str := string(data[:])
    end := strings.Index(str, FILE_NAME_WRAPPER_END)
    fileName := str[begin : end - 1]
    end += len(FILE_NAME_WRAPPER_END)
    path := getFilePath(fileName)
    homeDir := getSfttHomeDir()
    absPath := homeDir + path
    file, err := os.OpenFile(absPath, os.O_WRONLY | os.O_CREATE, 0666)
    common.CheckError(err)
    writer := bufio.NewWriter(file)
    writer.WriteString(str[end:])
    writer.Flush()
}
func DecodeDirContent(data []byte) {

}