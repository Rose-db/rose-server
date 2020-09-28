package main

import (
	"fmt"
	"strings"
	"sync"
)

type DbReadResult struct {
	Idx uint
	Id string
	Result string
}

/**
	A database is built with an Database.InternalDb that hold the values of
	the database in a key value map. Key is a uint while value is a []uint8 ([]byte).

	Individual indexes are stored in Database.IdLookupMap as a key value pair
	of user supplied key as a string and the index as uint.

    Database.InternalDb is growing in size as more values are stored. These size
	increaments can be called blocks.
	Blocks can hold up to 3000 indexes (value). When the reach max size, a new map
	is created with the same size.
 */
type Database struct {
	InternalDb map[uint]*[3000]*[]uint8
	// map of user supplied ids to InternalDb indexes
	// IdLookupMap::string -> idx::uint -> InternalDb[idx] -> []uint8
	IdLookupMap map[string]uint
	IdFactory *IdFactory
	RWMutex *sync.RWMutex

	CurrMapIdx uint
}

func NewDatabase() *Database {
	d := &Database{}

	d.InternalDb = make(map[uint]*[3000]*[]uint8)
	d.InternalDb[0] = &[3000]*[]uint8{}
	d.RWMutex = &sync.RWMutex{}

	d.IdLookupMap = make(map[string]uint)

	m := NewIdFactory()

	d.IdFactory = m
	d.CurrMapIdx = 0

	return d
}

/**
	- A RW lock is acquired
	- New uint idx is generated by IdFactory
 	- idx is stored in Database.IdLookupMap
	- a check is made for the current block
		- if the block does not exist, it is created
	- the value is stored in the block with its index
*/
func (d *Database) Insert(id string, v *[]uint8) uint {
	var idx uint
	var m *[3000]*[]uint8
	var computedIdx uint

	d.RWMutex.Lock()

	// r/w operation
	idx = d.IdFactory.Next()

	m, ok := d.InternalDb[d.CurrMapIdx]

	if !ok {
		m = &[3000]*[]uint8{}
		d.InternalDb[d.CurrMapIdx] = m
	}

	// r operation
	d.IdLookupMap[id] = idx
	m[idx] = v

	computedIdx = idx + (d.CurrMapIdx * 3000)

	if idx == 2999 {
		d.CurrMapIdx++
	}
	
	d.RWMutex.Unlock()

	return computedIdx
}

func (d *Database) Delete(id string) {

}

func (d *Database) Read(id string) (*DbReadResult, *DbReadError) {
	var idx uint
	var m *[3000]*[]uint8
	var mapId uint = 0
	var b *[]uint8

	idx, ok := d.IdLookupMap[id]

	if !ok {
		return nil, &DbReadError{
			Code:    InvalidReadErrorCode,
			Message: fmt.Sprintf("Invalid read operation. ID %s not exists", id),
		}
	}

	mapId = idx / 3000
	// get the map where the id value is
	m = d.InternalDb[mapId]

	// get the value of id
	b = m[idx]

	var sb strings.Builder
	sb.Grow(len(*b))

	for _, p := range *b {
		sb.WriteByte(p)
	}

	return &DbReadResult{
		Idx:    idx,
		Id:     id,
		Result: sb.String(),
	}, nil
}
