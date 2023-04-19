package main

import "fmt"

func Division(a int, b int) int {
	if a < b {
		return 0
	}
	return 1 + Division(a-b, b)
}

func main() {
	fmt.Println(Division(25, 3))
}
