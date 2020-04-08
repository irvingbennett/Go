package smallest

// Return prime factors in increasing order

import (
	"testing"
)

func TestSmallest(t *testing.T) {
	for _, test := range tests {
		actual := Smallest(test.input)
		if actual != test.expected {
			t.Fatalf("FAIL %s\nSmallest (%d) = %#v;\nexpected %#v",
				test.description, test.input,
				actual, test.expected)
		}
		t.Logf("PASS %s", test.description)
	}
}

/*
func BenchmarkPrimeFactors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Factors(test.input)
		}
	}
}
*/
