package main

import (
	"testing"
)

func TestIdGenerationWithinRange(t *testing.T) {
	var fac *IdMapFactory
	var curr int = 0

	fac = &IdMapFactory{}

	fac.Init()

	for {
		if curr == 10000000 {
			break
		}

		id := fac.Next()

		if id != curr {
			t.Errorf("TestIdGeneration: mismatched ids, expected %d, got %d", curr, id)

			return
		}

		curr++
	}
}

func BenchmarkIdFactory(b *testing.B) {
	var fac *IdMapFactory
	var curr int = 0

	fac = &IdMapFactory{}

	fac.Init()

	for n := 0; n < b.N; n++ {
		for {
			if curr == 100000000 {
				break
			}

			fac.Next()

			curr++
		}
	}
}