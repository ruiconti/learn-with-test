package iteration

import "testing"

func assertExpected(t *testing.T, expected, got string) {
	t.Helper()
	if expected != got {
		t.Errorf("Expected %q but got %q", expected, got)
	}
}

func TestRepeat(t *testing.T) {

	t.Run("Run without specifying how many iterations", func(t *testing.T) {
		repeated := Repeat("r", 5)
		expected := "rrrrr"
		assertExpected(t, expected, repeated)
	})

	t.Run("Run specifying how many iterations", func(t *testing.T) {
		repeated := Repeat("r", 15)
		expected := "rrrrrrrrrrrrrrr"
		assertExpected(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("r", 10)
	}
}

/*	Benchmark outputs:
	BenchmarkRepeat-4        6971016               178 ns/op
	which means that test ran 6971016 times averaging 178ns to run a single Repeat call
*/
