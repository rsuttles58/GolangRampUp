package diffsquares

//SquareOfSum receives a number and returns the Square of the Sum of all numbers from 0 to that number
func SquareOfSum(n int) int {

	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

//SumOfSquares receives a number and returns the Sum of Squares of each number from 0 to the input number
func SumOfSquares(n int) int {

	return (n * (n + 1) * (2*n + 1)) / 6
}

func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
