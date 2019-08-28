package pluck

import (
	"bytes"
	"html"

	"github.com/schollz/pluck/pluck/striphtml"
)

// Filter is a function type that takes a []byte and returns an []byte
type Filter = func(b []byte) []byte

// Sanitize sanitizes the input for common html characters
func Sanitize(b []byte) []byte {
	b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
	b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	b = []byte(striphtml.StripTags(html.UnescapeString(string(b))))
	return b
}

// TrimSpace trims the leading and trailing whitespace
func TrimSpace(b []byte) []byte {
	return bytes.TrimSpace(b)
}
