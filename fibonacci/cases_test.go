package fibonacci

// Source: exercism/problem-specifications
// Commit: bd2d4d9 sum-of-multiples: the factor 0 does not affect the sum of multiples of other factors
// Problem Specifications Version: 1.5.0

var varTests = []struct {
	n int
	f int
}{
	{0, 0},   // no multiples within limit
	{1, 1},   // one factor has multiples within limit
	{2, 1},   // more than one multiple within limit
	{3, 2},   // more than one factor with multiples within limit
	{4, 3},   // each multiple is only counted once
	{5, 5},   // 5
	{10, 55}, // a much larger limit
}
