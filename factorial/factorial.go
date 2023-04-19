package main

func Factorial(a int) int {
	if a < 2 {
		return 1
	}
	return a * Factorial(a-1)
}
