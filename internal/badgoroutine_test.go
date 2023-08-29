package profiling

import (
	"testing"
)

// go test -bench=BenchmarkNotEnoughChannels -blockprofile blockprofile.out
// go tool pprof blockprofile.out
// (pprof) tree
// (pprof) web
// (pprof) quit
func BenchmarkNotEnoughChannels(b *testing.B) {
	NewBadGoRoutine().NotEnoughChannels()
}
