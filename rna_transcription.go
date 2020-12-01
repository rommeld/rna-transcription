/*
Package strand changes DNA nucleotides for RNA.
*/
package strand

//ToRNA receives a DNA strand and returns its RNA complement.
func ToRNA(dna string) (solution string) {
	for _, nuc := range dna {
		if nuc == 'G' {
			solution += "C"
		} else if nuc == 'C' {
			solution += "G"
		} else if nuc == 'T' {
			solution += "A"
		} else if nuc == 'A' {
			solution += "U"
		}
	}
	return solution
}
