package profiling

import (
	"testing"
)

// go test -bench=BenchmarkTightLoop -cpuprofile cpuprofile.out
// go tool pprof -http=:8080 cpuprofile.out
// OR
// go test -trace traceprofile.out -bench=BenchmarkTightLoop
// go tool trace traceprofile.out
func BenchmarkTightLoop(b *testing.B) {
	NewBadCpu().TightLoop()
}
