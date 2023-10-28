package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	fmt.Scanln(&a)
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			fmt.Println("不是素数")
			return
		}
	}
	fmt.Println("是素数")
}
