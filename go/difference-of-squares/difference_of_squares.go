package diffsquares

const firstNonZeroNaturalNumber = 0

func SquareOfSum(n int) int {
	var squareOfSum int

	for i := 1; i <= n; i++ {
		squareOfSum += i
	}

	return squareOfSum * squareOfSum
}

func SumOfSquares(n int) int {
	var sumOfSquares int

	for i := 1; i <= n; i++ {
		sumOfSquares += i * i
	}

	return sumOfSquares
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
