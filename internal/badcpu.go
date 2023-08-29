package profiling

import "math/rand"

type BadCpu struct {
}

func NewBadCpu() *BadCpu {
	return &BadCpu{}
}

func (c *BadCpu) TightLoop() {
	for a := 0; a < 10000000; a++ {
		c.internalTightLoop()
	}
}

func (c *BadCpu) internalTightLoop() {
	rand.Int63()
}
