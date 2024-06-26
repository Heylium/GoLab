package benchmark

import (
	"bytes"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatStringByAdd(t *testing.T) {
	assert := assert2.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal("12345", ret)
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert2.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	var buf bytes.Buffer
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	assert.Equal("12345", buf.String())
}

/*
*
benchmark cmd:
win:
go test -bench="."
unix:
go test -bench=.

benchmark memory:
go test -bench=. -benchmem
*/
func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
