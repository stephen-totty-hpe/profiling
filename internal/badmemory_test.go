package profiling

import (
	"testing"
)

// go test -bench=BenchmarkContinuallyAddMore -memprofile memprofile.out
// go tool pprof -http=:8080 memprofile.out
// OR
// go test -trace traceprofile.out -bench=BenchmarkTightLoop
// go tool trace traceprofile.out
func BenchmarkContinuallyAddMore(b *testing.B) {
	b.ReportAllocs()
	NewBadMemory().ContinuallyAddMore()
}

func TestContinuallyAddMore(t *testing.T) {
	NewBadMemory().ContinuallyAddMore()
}
