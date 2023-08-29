package profiling

import (
	"time"
)

// https://go101.org/article/concurrent-common-mistakes.html
type BadGoRoutine struct {
}

func NewBadGoRoutine() *BadGoRoutine {
	return &BadGoRoutine{}
}

func (r *BadGoRoutine) NotEnoughChannels() int {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			r.internalNotEnoughChannels(i, c)
		}()
	}
	time.Sleep(time.Second * 3)
	return <-c
}

func (r *BadGoRoutine) internalNotEnoughChannels(j int, myChan chan int) {
	// do some other non interesting stuff
	x := 0
	for a := 0; a < 10000; a++ {
		x++
	}

	r.coreNotEnoughChannels(j, myChan)

	for b := 0; b < 10000; b++ {
		x++
	}
}

func (r *BadGoRoutine) coreNotEnoughChannels(j int, myChan chan int) {
	myChan <- j // 4 goroutines will hang here.
}
