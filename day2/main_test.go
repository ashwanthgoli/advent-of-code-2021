package main

import "testing"

func BenchmarkNavigateP1V1(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		navigateP1(testInput)
	}
}

var testInput = `
forward 9
down 3
down 8
forward 2
up 3
forward 5
up 8
down 2
down 5
up 7
down 9
forward 4
up 5
down 9
forward 2
forward 2
forward 8
down 6
`
