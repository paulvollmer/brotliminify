package brotliminify

import (
	"bytes"
	"io"
	"regexp"

	"github.com/google/brotli/go/cbrotli"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"
)

// Encode minify and encode a given source and return the bytes
func Encode(source io.Reader) ([]byte, error) {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)

	var b bytes.Buffer
	err := m.Minify("text/html", &b, source)
	if err != nil {
		return nil, err
	}
	b2, err := cbrotli.Encode(b.Bytes(), cbrotli.WriterOptions{Quality: 11})
	if err != nil {
		return nil, err
	}
	return b2, nil
}

// Decode a given source and return the bytes
func Decode(source []byte) ([]byte, error) {
	return cbrotli.Decode(source)
}
