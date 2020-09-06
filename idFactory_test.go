package main

import (
	"testing"
)

func TestIdGenerationWithinRange(t *testing.T) {
	var fac *IdFactory
	var iterations int = 0
	var currId uint = 0

	fac = &IdFactory{}

	fac.Init()

	for {
		if iterations == 10000000 {
			break
		}

		id := fac.Next()

		if id < 0 || id > 2999 {
			t.Errorf("Invalid id produced by the IdFactory. Expected a value between 0 and 3000, got %d", id);

			return
		}

		if currId != id {
			t.Errorf("Counted ins are not equal. Got %d, expected %d", id, currId);

			return
		}

		currId++

		iterations++

		if currId > 2999 {
			currId = 0
		}
	}
}