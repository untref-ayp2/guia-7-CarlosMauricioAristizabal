package main

import "fmt"

func Producto(a int, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	return a + Producto(a, b-1)
}

func main() {
	fmt.Println(Producto(5, 3))
}
