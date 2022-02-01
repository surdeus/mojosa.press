package md

import(
	bf "github.com/russross/blackfriday/v2"
)

func
Process(s []byte) []byte {
	return bf.Run(s)
}
