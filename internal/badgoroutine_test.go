package profiling

import (
	"testing"
)

// go test -bench=BenchmarkNotEnoughChannels -blockprofile blockprofile.out
// go tool pprof -http=:8080 blockprofile.out
// OR
// go test -trace traceprofile.out -bench=BenchmarkTightLoop
// go tool trace traceprofile.out
func BenchmarkNotEnoughChannels(b *testing.B) {
	NewBadGoRoutine().NotEnoughChannels()
}

func TestNotEnoughChannels(t *testing.T) {
	NewBadGoRoutine().NotEnoughChannels()
}
