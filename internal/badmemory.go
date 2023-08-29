package profiling

import "time"

// https://go101.org/article/memory-leaking.html
type BadMemory struct {
	bytes1 []byte
	bytes2 []byte
}

func NewBadMemory() *BadMemory {
	return &BadMemory{
		bytes1: []byte{},
		bytes2: []byte{},
	}
}

func (m *BadMemory) ContinuallyAddMore() {
	for a := 0; a < 5; a++ {
		m.AddMore1(100000)

		m.AddMore2(2000000)
		time.Sleep(time.Second * 1)
	}
}

func (m *BadMemory) AddMore1(numBytes int) {
	localBytes := make([]byte, numBytes)
	m.bytes1 = append(m.bytes1, localBytes...)
}

func (m *BadMemory) AddMore2(numBytes int) {
	localBytes := make([]byte, numBytes)
	m.bytes2 = append(m.bytes2, localBytes...)
}
