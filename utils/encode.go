package utils

//import (
//	iconv "github.com/djimenez/iconv-go"
//)

var (
// CVtolocal *iconv.Converter
// CVtoutf8  *iconv.Converter
)

func InitConverter(local string) bool {
	return true
}

func Utf8ToLocal(str string) (b string, err error) {
	return str, err
}

func LocalToUtf8(str string) (b string, err error) {
	return str, err
}
