package main

import (
	"fmt"
	"strings"
)

type Database struct {
	InternalDb map[int]*[10000]*[]byte
	IdLookupMap map[string]int
	IdMapFactory IdMapFactory
}

func (d *Database) Init() {
	d.InternalDb = make(map[int]*[10000]*[]byte)
	d.IdLookupMap = make(map[string]int)

	m := &IdMapFactory{}
	m.Init()

	d.IdMapFactory = *m
}

func (d *Database) Insert(id string, v *[]byte) {

}

func (d *Database) Delete(id int) {

}

func (d *Database) Read(id int) {
	b := []byte{}
	s := "test"
	var sb strings.Builder
	sb.Grow(len(s))

	for _, p := range b {
		sb.WriteByte(p)
	}

	s = sb.String()   // no copying
	fmt.Println(s)
}
