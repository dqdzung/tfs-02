package myPackage

import (
	"fmt"
)

// Find min max from number slice
func FindMinMax(sl []int) (int, int) {
	var max int = sl[0]
	var min int = sl[0]
	for _, v := range sl {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	return min, max
}

// Get average value from number slice
func GetAverage(sl []int) int {
	sum := int(0)
	for _, v := range sl {
		sum += v
	}
	return sum / len(sl)
}

// Check if a given number is prime
func IsPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

// Check if a given number exists in slice
func DoExist(n int, sl []int) bool {
	for k, v := range sl {
		if v == n {
			fmt.Printf("Found number %v at position %v\n", v, k)
			return true
		}
	}
	fmt.Printf("Number %v not found in slice\n", n)
	return false
}
