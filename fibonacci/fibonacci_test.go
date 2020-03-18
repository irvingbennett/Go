package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	for _, test := range varTests {
		s := Fibonacci(test.n)
		if s != test.f {
			t.Fatalf("Fibonacci of %d returned %d, want %d.",
				test.n, s, test.f)
		}
	}
}

/*
func BenchmarkSumMultiples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range varTests {
			SumMultiples(test.limit, test.divisors...)
		}
	}
}
*/
