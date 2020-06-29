package hamming

import "errors"

// Distance calculate hamming disatance
func Distance(a, b string) (int, error) {
	var firstStringLength = len(a)
	var secondStringLength = len(b)

	if firstStringLength != secondStringLength {
		return 0, errors.New("string length are not equal")
	}

	var incorrectGenes = 0
	for i := 0; i < firstStringLength; i++ {
		if a[i] != b[i] {
			incorrectGenes++
		}
	}
	return incorrectGenes, nil
}
