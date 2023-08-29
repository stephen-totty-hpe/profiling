package profiling

import "time"

type RunMain struct {
	badCpu    *BadCpu
	badMemory *BadMemory
}

func NewRunMain() *RunMain {
	return &RunMain{
		badCpu:    NewBadCpu(),
		badMemory: NewBadMemory(),
	}
}

func (r *RunMain) Run() {
	for i := 0; i < 10000; i++ {
		r.badCpu.internalTightLoop()
		r.badMemory.AddMore1(100)
		time.Sleep(time.Millisecond * 10)
	}
}
