package main

import (
	"fmt"
	"math"
)

func isPrime(a int) bool {
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var a int
	fmt.Scanln(&a)

	if isPrime(a) {
		fmt.Printf("%d 是素数\n", a)
	} else {
		fmt.Printf("%d 不是素数\n", a)
	}
}
