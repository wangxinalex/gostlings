// Concept: benchmark setup and allocation reporting
// Task: complete the benchmark, then run it with `go test -bench=. -benchmem`
// Expected output: PASS and a benchmark result (run with `go test -bench=. -benchmem ./exercises/14_testing/testing11`)
// Hint: prepare input before timing starts. Call b.ReportAllocs(), then b.ResetTimer(),
//       and call Count(input) inside a loop from 0 to b.N. The benchmark does not
//       need to print anything; the testing package reports timing and allocations.

package lines

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	if got := Count("a\nb\n"); got != 2 {
		t.Errorf("Count() = %d, want 2", got)
	}
}

func BenchmarkCount(b *testing.B) {
	input := strings.Repeat("line\n", 1000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Count(input)
	}
}
