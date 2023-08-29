package profiling

import (
	"testing"
)

// go test -bench=BenchmarkContinuallyAddMore -memprofile memprofile.out
// go tool pprof memprofile.out
// (pprof) tree
// (pprof) web
// (pprof) quit
func BenchmarkContinuallyAddMore(b *testing.B) {
	NewBadMemory().ContinuallyAddMore()
}
