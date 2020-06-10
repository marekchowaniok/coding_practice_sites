package diffsquares

//SquareOfSum first calculate sum and then square
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

//SumOfSquares first calculate square and then sum
func SumOfSquares(n int) int {
	sum := (n * (n + 1) * (2*n + 1)) / 6
	return sum
}

//Difference function to calculate diff between SquareOfSum and SumOfSquares
func Difference(n int) interface{} {
	sum := ((n - 1) * n * (n + 1) * (3*n + 2)) / 12
	return sum
}
