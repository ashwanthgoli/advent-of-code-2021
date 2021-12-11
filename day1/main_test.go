package main

import (
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
199
200
208
210
200
207
240
269
260
263
`)

func BenchmarkDiveV1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		divep1v1(testInput)
	}
}

func BenchmarkDiveV2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		divep1v2(testInput)
	}
}

func BenchmarkDiveBV1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		divep2v1(testInput)
	}
}

func BenchmarkDiveBV2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		divep2v2(testInput)
	}
}
