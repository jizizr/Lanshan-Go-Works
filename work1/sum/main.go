package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println(sum(a, b))
}
