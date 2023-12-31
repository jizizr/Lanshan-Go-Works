package main

import (
	"fmt"
	"math"
)

func area(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

func main() {
	fmt.Print("圆的半径为：")
	var radius float64
	fmt.Scanln(&radius)
	fmt.Println("圆的面积为：", area(radius))
}
