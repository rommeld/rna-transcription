/*
Package strand changes DNA nucleotides for RNA.
*/
package strand

import "bytes"

var transcript = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

//ToRNA receives a DNA strand and returns its RNA complement.
func ToRNA(dna string) (solution string) {
	var buf bytes.Buffer

	/*
		When working with a if conditionals you do not need to define a transcript
		or buf. Therefore bytes has not to be imported.
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
	*/

	for _, nuc := range dna {
		buf.WriteString(transcript[nuc])
	}
	return buf.String()
}
