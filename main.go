package main

func main() {
	var a *AppController
	var m *Metadata

	m = &Metadata{
		Method: "insert",
		Data: &[]byte{},
	}

	a = &AppController{}

	a.Run(m)
}