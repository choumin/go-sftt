package decode

import (
	"../fileops"
)

func decipher(data []byte) (ret []byte) {
	ret = data
	return
}
func Decode(data []byte) {
	// todo, use decipher algo to decode data
	data = decipher(data)
	fileType := fileops.ParseFileType(data)
	if fileType == fileops.TYPE_FILE {
		fileops.DecodeFileContent(data)
	} else {
		fileops.DecodeDirContent(data)
	}
}