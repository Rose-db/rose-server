package main

type IdMapFactory struct {
	IdMap map[int]*[100]uint8
	CurrMap int
	CurrIdx uint8
}

func (m *IdMapFactory) Init() {
	m.IdMap = make(map[int]*[100]uint8)
	m.IdMap[0] = &[100]uint8{}

	m.CurrMap = 0
	m.CurrIdx = 0
}

func (m *IdMapFactory) Next() int {
	var a *[100]uint8

	a = m.IdMap[m.CurrMap]

	for i := range a {
		if a[i] == 0 {
			a[i] = 1

			if m.CurrMap == 0 {
				return i
			}

			return i + (m.CurrMap * 100)
		}
	}

/*	if m.CurrIdx < 100 {
		// mark as taken
		a[m.CurrIdx] = 1

		id = int(a[m.CurrIdx])
		m.CurrIdx++

		if m.CurrMap != 0 {
			return id + (m.CurrMap * 100)
		}

		return id
	}*/


	// safe to assume that the current block is full,
	// so we must allocate the next block
	m.CurrMap++
	m.IdMap[m.CurrMap] = &[100]uint8{}

	return m.Next()
}
