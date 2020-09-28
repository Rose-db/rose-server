package main

/**
	Creates the next id to be generated. It only generates id <= IdFactory.Max.
	How these ids are represented in the database is handled by the Database
	struct. After it reaches IdFactory.Max, it resets to 0.
 */
type IdFactory struct {
	Max uint16
	CurrIdx uint16
}

func NewIdFactory() *IdFactory {
	return &IdFactory{
		Max: 2999,
		CurrIdx: 0,
	}
}

func (m *IdFactory) Next() uint {
	if m.CurrIdx == 0 {
		m.CurrIdx++

		return uint(0)
	}

	var c uint16

	c = m.CurrIdx
	m.CurrIdx++

	if m.CurrIdx > 2999 {
		m.CurrIdx = 0
	}

	return uint(c)
}
