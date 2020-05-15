package fileops

import (
    "bufio"
    "bytes"
    "encoding/binary"
    "fmt"
    "github.com/choumin/sftt/internal/app/common"
    "io/ioutil"
    "os"
    "strings"
)

type MetaFileContent struct{
    FileType uint16
    FileNameLen uint8
    FileName [common.MAX_FILE_NAME_LEN]byte
}

const BUF_SIZE = 1024000
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
func getFileName(path string) (string) {
    index := strings.LastIndex(path, "/")
    if index == -1 {
        return path
    }
    return path[index + 1 : ]
}
func (obj *MetaFileContent)structToByteArray() ([]byte, error) {
    buf := new(bytes.Buffer)
    if err := binary.Write(buf, binary.LittleEndian, obj); err != nil {
        fmt.Println("write failed: ", err)
        return nil, err
    }
    return buf.Bytes(), nil
}
func GenerateMetaFileContent(path string) (data []byte) {
    mateFileContent := &MetaFileContent{}
    mateFileContent.FileType = common.FILE_TYPE_MAGIC_FILE
    tmp := []byte(getFileName(path))
    mateFileContent.FileNameLen = uint8(len(tmp))
    for i, byte := range tmp {
        if (i >= common.MAX_FILE_NAME_LEN) {
            break
        }
        mateFileContent.FileName[i] = byte
    }

    fmt.Println(mateFileContent.FileType)
    fmt.Println(mateFileContent.FileNameLen)
    fmt.Println(mateFileContent.FileName)
    //fmt.Print(len(mateFileContent.FileName))
    data, _ = mateFileContent.structToByteArray()
    //fmt.Println(data)
    //copy(mateFileContent.FileName, [common.MAX_FILE_NAME_LEN]byte(getFileName(path)))
    return
}

func FetchAllBytes(path string) (data []byte) {

    return
}