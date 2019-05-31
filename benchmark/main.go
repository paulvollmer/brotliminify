package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"

	"github.com/paulvollmer/brotliminify"
)

func main() {
	compressGzip(textHTML)
	compressBrotliminify(bytes.NewReader(textHTML))
}

func compressGzip(source []byte) {
	var b bytes.Buffer
	gz, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		panic(err)
	}
	if _, err := gz.Write(source); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	fmt.Println("GZIP SIZE:\t", len(b.Bytes()))
}

func compressBrotliminify(source io.Reader) {
	result, err := brotliminify.Encode(source)
	if err != nil {
		panic(err)
	}
	fmt.Println("BROTLIMIN SIZE:\t", len(result))
}
