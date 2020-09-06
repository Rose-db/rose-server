package main

import (
	"fmt"
	"testing"
)

// test string of 729 bytes
var testString string = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus orci magna, dapibus eget tristique rhoncus, ornare et leo. Cras pellentesque urna vitae varius volutpat. Integer vulputate pulvinar mattis. Cras varius sapien tempor posuere volutpat. Vivamus eu iaculis mauris. Donec accumsan, ante at ornare porttitor, ex nisl rutrum justo, eget convallis turpis sapien vulputate est. Sed blandit sapien quis ex ornare, a cursus nisl tincidunt. Phasellus posuere pellentesque felis a vulputate. Praesent nec ante non nisi bibendum rutrum vulputate ut tortor. Donec nibh arcu, bibendum vel lacinia sit amet, hendrerit commodo quam. Suspendisse elit eros, facilisis non dictum eget, efficitur ac ligula. Suspendisse id felis lacus. "

func benchmarkInsert(i int, a *AppController, m *Metadata, b *testing.B) {
	for c := 0; c < i; c++ {
		_, _ = a.Run(m)
	}
}

func BenchmarkInsertHundred(b *testing.B) {
	var s []byte
	var a *AppController
	var m *Metadata

	s = []byte(testString)

	a = &AppController{}
	a.Init()

	for n := 0; n < b.N; n++ {
		m = &Metadata{
			Method: InsertMethodType,
			Data: &s,
			Id: fmt.Sprintf("id-%d", n),
		}

		benchmarkInsert(100, a, m, b)
	}
}

func BenchmarkInsertThousand(b *testing.B) {
	var s []byte
	var a *AppController
	var m *Metadata

	s = []byte(testString)

	a = &AppController{}
	a.Init()

	for n := 0; n < b.N; n++ {
		m = &Metadata{
			Method: InsertMethodType,
			Data: &s,
			Id: fmt.Sprintf("id-%d", n),
		}

		benchmarkInsert(1000, a, m, b)
	}
}

func BenchmarkInsertTenThousand(b *testing.B) {
	var s []byte
	var a *AppController
	var m *Metadata

	s = []byte(testString)

	a = &AppController{}
	a.Init()

	for n := 0; n < b.N; n++ {
		m = &Metadata{
			Method: InsertMethodType,
			Data: &s,
			Id: fmt.Sprintf("id-%d", n),
		}

		benchmarkInsert(10000, a, m, b)
	}
}

func BenchmarkInsertHundredThousand(b *testing.B) {
	var s []byte
	var a *AppController
	var m *Metadata

	s = []byte(testString)

	a = &AppController{}
	a.Init()

	for n := 0; n < b.N; n++ {
		m = &Metadata{
			Method: InsertMethodType,
			Data: &s,
			Id: fmt.Sprintf("id-%d", n),
		}

		benchmarkInsert(100000, a, m, b)
	}
}

func BenchmarkInsertMillion(b *testing.B) {
	var s []byte
	var a *AppController
	var m *Metadata

	s = []byte(testString)

	a = &AppController{}
	a.Init()

	for n := 0; n < b.N; n++ {
		m = &Metadata{
			Method: InsertMethodType,
			Data: &s,
			Id: fmt.Sprintf("id-%d", n),
		}

		benchmarkInsert(1000000, a, m, b)
	}
}

func BenchmarkInsertHundredMillion(b *testing.B) {
	var s []byte
	var a *AppController
	var m *Metadata

	s = []byte(testString)

	a = &AppController{}
	a.Init()

	for n := 0; n < b.N; n++ {
		m = &Metadata{
			Method: InsertMethodType,
			Data: &s,
			Id: fmt.Sprintf("id-%d", n),
		}

		benchmarkInsert(100000000, a, m, b)
	}
}
