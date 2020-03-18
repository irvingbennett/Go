package smallest

// Source: exercism/problem-specifications
// Commit: d928874 prime-factors: apply "input" policy
// Problem Specifications Version: 1.1.0

var tests = []struct {
	description string
	input       int64
	expected    int64
}{

	{
		"no factors",
		1,
		1,
	},
	{
		"prime number",
		2,
		2,
	},
	{
		"square of a prime",
		3,
		6,
	},
	{
		"cube of a prime",
		4,
		12,
	},

	{
		"range of twenty",
		6,
		60,
	},
	{
		"range of ten",
		10,
		2520,
	},
	{
		"range of twenty",
		20,
		232792560,
	},
}
