package encode

func crypto(data []byte) (code []byte) {
	code = data
	return
}
//func Encode(path string) (data []byte) {
//	if fileops.IsFile(path) {
//		data = fileops.EncodeFileContent(path)
//	} else {
//		data = fileops.EncodeDirContent(path)
//	}
//	// todo, use crypto algo to encode data
//	data = crypto(data)
//	return
//}

func EncodeBytes(input []byte) (output []byte) {
	output = crypto(input)
	return
}