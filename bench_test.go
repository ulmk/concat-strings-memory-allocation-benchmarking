package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// go test -benchmem -bench=. bench_test.go
func BenchmarkConcat(b *testing.B) {
	str1 := "Hello"
	str2 := "World"

	b.Run("PlusOperator", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := str1 + " " + str2
			_ = result
		}
	})

	b.Run("Sprintf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := fmt.Sprintf("%s %s", str1, str2)
			_ = result
		}
	})

	b.Run("StringBuilder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var b strings.Builder
			b.WriteString(str1)
			b.WriteString(" ")
			b.WriteString(str2)
			result := b.String()
			_ = result
		}
	})

	b.Run("ByteBuffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			buf.WriteString(str1)
			buf.WriteString(" ")
			buf.WriteString(str2)
			result := buf.String()
			_ = result
		}
	})

	b.Run("StringsJoin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			strs := []string{str1, str2}
			result := strings.Join(strs, "")
			_ = result
		}
	})
}
