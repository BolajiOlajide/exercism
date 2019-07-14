package hamming

import "errors"

// Distance calculate hamming disatance
func Distance(a, b string) (int, error) {
	var firstStringLength = len(a)
	var secondStringLength = len(b)
	var incorrectGenes = 0

	if firstStringLength != secondStringLength {
		return 0, errors.New("String length are not equal")
	}

	for i := 0; i < firstStringLength; i++ {
		if a[i] != b[i] {
			incorrectGenes++
		}
	}
	return incorrectGenes, nil
}
