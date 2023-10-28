package main

import (
	"fmt"
	"math/rand"
	"time"
)

func binarySearch(target int, low int, high int) (int, bool) {
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		if mid == target {
			return mid, true
		} else if mid < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, false
}

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)
	var target int = randomGenerator.Intn(100) + 1

	if num, flag := binarySearch(target, 1, 100); flag {
		fmt.Printf("找到了数字，是%d\n", num)
	} else {
		fmt.Println("没有找到数字。是不是不在1-100之间？")
	}
}
