package main

import (
	"strings"
)

type Database struct {
	InternalDb map[uint]*[3000]*[]uint8
	IdLookupMap map[string]uint
	IdFactory *IdFactory

	CurrMapIdx uint
}

func (d *Database) Init() {
	d.InternalDb = make(map[uint]*[3000]*[]uint8)
	d.InternalDb[0] = &[3000]*[]uint8{}

	d.IdLookupMap = make(map[string]uint)

	m := &IdFactory{}
	m.Init()

	d.IdFactory = m
	d.CurrMapIdx = 0
}

func (d *Database) Insert(id string, v *[]uint8) uint {
	var idx uint
	var m *[3000]*[]uint8
	var computedIdx uint

	idx = d.IdFactory.Next()

	d.IdLookupMap[id] = idx

	m = d.InternalDb[d.CurrMapIdx]

	if m == nil {
		m = &[3000]*[]uint8{}
	}

	m[idx] = v

	d.InternalDb[d.CurrMapIdx] = m
	computedIdx = idx + (d.CurrMapIdx * 3000)

	if idx == 2999 {
		d.CurrMapIdx++
	}

	return computedIdx
}

func (d *Database) Delete(id string) {

}

func (d *Database) Read(id string) string {
	var Id uint
	var m *[3000]*[]uint8
	var mapId uint = 0
	var b *[]uint8

	Id = d.IdLookupMap[id]
	mapId = Id / 3000
	m = d.InternalDb[mapId]

	b = m[Id]

	var sb strings.Builder
	sb.Grow(len(*b))

	for _, p := range *b {
		sb.WriteByte(p)
	}

	return sb.String()
}
