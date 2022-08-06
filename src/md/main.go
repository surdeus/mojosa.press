package md

import(
	bf "github.com/russross/blackfriday/v2"
	"strings"
)

func
Process(s []byte) []byte {
	s = []byte(strings.Replace(string(s), "\r", "", -1))
	return bf.Run(s)
}
