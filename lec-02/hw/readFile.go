package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Get string slice from file
func readFile() []string {
	data, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return []string{}
	}
	content := string(data)
	s := strings.Split(content, " ")
	return s
}

// Get number slice from string
func getNumbers() []int {
	s := readFile()
	results := []int{}
	for _, v := range s {
		val, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		results = append(results, val)
	}
	return results
}


// Find min max from number slice
func findMinMax(sl []int) (int, int) {
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
func getAverage(sl []int) int {
	sum := int(0)
	for _, v := range sl {
		sum += v
	}
	return sum / len(sl)
}

// Check if a given number is prime
func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

// Check if a given number exists in slice
func doExist(n int, sl []int) bool {
	for k, v := range sl {
		if v == n {
			fmt.Printf("Found number %v at position %v\n", v, k)
			return true
		}
	}
	fmt.Printf("Number %v not found in slice\n", n)
	return false
}


func main() {
	nums := getNumbers()
	fmt.Println("Nums:", nums)

	min, max := findMinMax(nums)
	fmt.Printf("Min: %v\nMax: %v\n", min, max)

	avg := getAverage(nums)
	fmt.Println("Average:", avg)

	primes := []int{}
	for _, v := range nums {
		isPrime := isPrime(v)
		if isPrime {
			primes = append(primes, v)
		}
	}
	fmt.Println("Primes:", primes)

	doExist(42, nums)
	doExist(69, nums)
}