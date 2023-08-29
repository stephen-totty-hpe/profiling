package profiling

import (
	"testing"
)

// go test -bench=BenchmarkTightLoop -cpuprofile cpuprofile.out
// go tool pprof cpuprofile.out
// (pprof) tree
// (pprof) web
// (pprof) quit
func BenchmarkTightLoop(b *testing.B) {
	NewBadCpu().TightLoop()
}
