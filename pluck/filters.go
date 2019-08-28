package pluck

import (
	"bytes"
	"html"

	"github.com/schollz/pluck/pluck/striphtml"
)

type Filter = func(b []byte) []byte

func Sanitize(b []byte) []byte {
	b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
	b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	b = []byte(striphtml.StripTags(html.UnescapeString(string(b))))
	return b
}

func TrimSpace(b []byte) []byte {
	return bytes.TrimSpace(b)
}
