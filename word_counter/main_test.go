package main

import (
	"bytes"
	"testing"
)

func makeInput(words int) []byte {
	var b bytes.Buffer
	for i := 0; i < words; i++ {
		b.WriteString("word ")
		if i%20 == 0 {
			b.WriteByte('\n')
		}
	}
	return b.Bytes()
}

var data = makeInput(200_000)

func BenchmarkWordCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(data)
		_, _ = countWordsStream(r)
	}
}

func BenchmarkWordCountNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(data)
		_, _ = countWordsStreamNative(r)
	}
}
