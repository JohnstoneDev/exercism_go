package strand

func ToRNA(dna string) string {
	rna := []rune(dna)

	for i, letter := range dna {
		switch letter {
		case 'A':
			rna[i] = 'U'
		case 'C':
			rna[i] = 'G'
		case 'G':
			rna[i] = 'C'
		case 'T':
			rna[i] = 'A'
		}
	}

	return string(rna)
}
