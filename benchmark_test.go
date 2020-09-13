package main

import (
	"fmt"
	"testing"
)

// test string of 2,981 bytes
var testString string = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus ac pretium nunc. Pellentesque egestas rutrum neque vitae pellentesque. Sed quis dolor ut velit congue aliquam a vitae nulla. Morbi maximus metus quis neque commodo posuere. Nunc sit amet luctus sapien. Donec pharetra, urna non commodo posuere, purus nisl rutrum justo, eu eleifend dolor turpis cursus libero. Fusce vel velit et neque laoreet scelerisque eu non ex. Proin tempor viverra eleifend. Phasellus aliquam, massa a tincidunt maximus, sapien augue commodo diam, quis scelerisque eros purus nec lectus. Nulla varius condimentum erat congue venenatis. Aenean vel mauris cursus, feugiat lorem a, pellentesque lorem. Nam diam dolor, semper non augue sit amet, posuere tempor nunc. Aliquam elit sapien, placerat vel eros dignissim, fermentum eleifend dolor. Aenean auctor quis ex scelerisque varius. Nulla aliquam dapibus viverra. Suspendisse sit amet metus imperdiet odio porttitor imperdiet.\n\nProin id nulla rutrum, bibendum eros vel, interdum mauris. Pellentesque a mattis elit. Maecenas sodales magna in nunc auctor, vel rhoncus urna elementum. Phasellus tristique dictum lorem, vel placerat urna sollicitudin non. Cras id tincidunt lorem, id rutrum nisl. Vestibulum sed egestas justo. Proin dui est, bibendum ut nulla a, dignissim rhoncus sapien. Maecenas in varius sem, in tristique quam.\n\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur finibus posuere turpis, id ullamcorper turpis laoreet auctor. Integer eu nulla vel odio tempus porta. Fusce vitae lorem ac erat mollis sollicitudin. Maecenas congue efficitur mollis. Etiam ac ex facilisis, imperdiet magna sed, pulvinar quam. Maecenas et ante leo.\n\nNulla malesuada tellus eu lorem malesuada vehicula. Nam vel vestibulum enim, in accumsan metus. Donec commodo, nisi in varius consectetur, felis sapien pellentesque tellus, nec consequat diam ante nec magna. Maecenas ut faucibus nisl. Fusce non nisl vitae risus aliquam aliquet nec in enim. Maecenas erat lacus, pharetra ac pharetra ac, aliquam nec magna. Aenean et quam elit. Suspendisse ornare volutpat odio vel tempus. Maecenas eget erat a est aliquet semper. Nunc tincidunt tincidunt ullamcorper. Donec sit amet velit pulvinar, ornare orci sed, sodales leo. Proin felis purus, maximus non ante eu, sagittis molestie justo.\n\nCras dapibus tellus leo, quis imperdiet orci pulvinar in. Quisque ultricies tellus non tincidunt porta. Morbi at neque id eros consectetur hendrerit ut id diam. Suspendisse fringilla lorem quis feugiat dapibus. Integer fermentum pulvinar ipsum id vulputate. Fusce tellus nunc, sagittis a tincidunt sit amet, posuere ac lectus. Suspendisse potenti. Proin feugiat erat justo, in ultrices leo elementum ut. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis quis blandit quam. Fusce sed turpis est. Nulla lacus magna, pulvinar nec auctor ac, interdum a nunc. Aenean in est leo. Integer blandit diam at tortor euismod, eget mattis lectus aliquet. "

func benchmarkDirectInsert(i int, a *AppController, m *Metadata) {
	for c := 0; c < i; c++ {
		_, _ = a.Run(m)
	}
}

func benchmarkServerInsert(i int, method string, data string, b *testing.B) {
	for c := 0; c < i; c++ {
		testSendRequest(method, fmt.Sprintf("id-%d", c), data, b)
	}
}

func BenchmarkDirectInsertHundred(b *testing.B) {
	var s []byte
	var a *AppController
	var m *Metadata

	a = testCreateController(testGetBenchmarkName(b))
	s = []byte(testString)

	for n := 0; n < b.N; n++ {
		m = &Metadata{
			Method: InsertMethodType,
			Data: &s,
			Id: fmt.Sprintf("id-%d", n),
		}

		benchmarkDirectInsert(100, a, m)
	}
}

func BenchmarkDirectInsertThousand(b *testing.B) {
	var s []byte
	var a *AppController
	var m *Metadata

	s = []byte(testString)
	a = testCreateController(testGetBenchmarkName(b))

	for n := 0; n < b.N; n++ {
		m = &Metadata{
			Method: InsertMethodType,
			Data: &s,
			Id: fmt.Sprintf("id-%d", n),
		}

		benchmarkDirectInsert(1000, a, m)
	}
}

func BenchmarkDirectInsertTenThousand(b *testing.B) {
	var s []byte
	var a *AppController
	var m *Metadata

	s = []byte(testString)
	a = testCreateController(testGetBenchmarkName(b))

	for n := 0; n < b.N; n++ {
		m = &Metadata{
			Method: InsertMethodType,
			Data: &s,
			Id: fmt.Sprintf("id-%d", n),
		}

		benchmarkDirectInsert(10000, a, m)
	}
}

func BenchmarkDirectInsertHundredThousand(b *testing.B) {
	var s []byte
	var a *AppController
	var m *Metadata

	s = []byte(testString)
	a = testCreateController(testGetBenchmarkName(b))

	for n := 0; n < b.N; n++ {
		m = &Metadata{
			Method: InsertMethodType,
			Data: &s,
			Id: fmt.Sprintf("id-%d", n),
		}

		benchmarkDirectInsert(100000, a, m)
	}
}

func BenchmarkDirectInsertMillion(b *testing.B) {
	var s []byte
	var a *AppController
	var m *Metadata

	s = []byte(testString)
	a = testCreateController(testGetBenchmarkName(b))

	for n := 0; n < b.N; n++ {
		m = &Metadata{
			Method: InsertMethodType,
			Data: &s,
			Id: fmt.Sprintf("id-%d", n),
		}

		benchmarkDirectInsert(1000000, a, m)
	}
}

func BenchmarkDirectInsertHundredMillion(b *testing.B) {
	b.Skip(fmt.Sprintf("Skip %s", testGetBenchmarkName(b)))

	var s []byte
	var a *AppController
	var m *Metadata

	s = []byte(testString)
	a = testCreateController(testGetBenchmarkName(b))

	for n := 0; n < b.N; n++ {
		m = &Metadata{
			Method: InsertMethodType,
			Data: &s,
			Id: fmt.Sprintf("id-%d", n),
		}

		benchmarkDirectInsert(100000000, a, m)
	}
}

/****************************************

FOR THESE BENCHMARKS, ROSE SERVER MUST BE STARTED

****************************************/
func BenchmarkServerInsertHundred(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkServerInsert(100, InsertMethodType, testString, b)
	}
}

func BenchmarkServerInsertThousand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkServerInsert(1000, InsertMethodType, testString, b)
	}
}

func BenchmarkServerInsertTenThousand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkServerInsert(10000, InsertMethodType, testString, b)
	}
}

func BenchmarkServerInsertHundredThousand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		benchmarkServerInsert(100000, InsertMethodType, testString, b)
	}
}

func BenchmarkServerInsertMillion(b *testing.B) {
	b.Skip(fmt.Sprintf("%s: Skipping... overly intensive", testGetBenchmarkName(b)))

	for n := 0; n < b.N; n++ {
		benchmarkServerInsert(1000000, InsertMethodType, testString, b)
	}
}
