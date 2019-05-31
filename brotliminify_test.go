package brotliminify

import (
	"testing"
	"strings"
)

func TestCompressHTML(t *testing.T) {
	testSource := `<div><h1>test</h1></div>`

	compressedResult, err := Encode(strings.NewReader(testSource))
	if err != nil {
		t.Error(err)
	}

	decompressedResult, err := Decode(compressedResult)
	if err != nil {
		t.Error(err)
	}
	if testSource != string(decompressedResult) {
		t.Error("decompressedResult not equal")
	}
}
