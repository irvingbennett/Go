// Package strand encode rna strands
package strand

// ToRNA returns the RNA of a DNA string
func ToRNA(dna string) (rna string) {
	t := map[string]byte{
		"G": 'C',
		"C": 'G',
		"T": 'A',
		"A": 'U',
	}

	r := make([]byte, 0)
	for _, x := range dna {
		r = append(r, t[string(x)])
	}
	rna = string(r)
	return
}
