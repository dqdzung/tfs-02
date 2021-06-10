package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"../myPackage"
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

func main() {
	nums := getNumbers()
	fmt.Println("Nums:", nums)

	min, max := myPackage.FindMinMax(nums)
	fmt.Printf("Min: %v\nMax: %v\n", min, max)

	avg := myPackage.GetAverage(nums)
	fmt.Println("Average:", avg)

	primes := []int{}
	for _, v := range nums {
		isPrime := myPackage.IsPrime(v)
		if isPrime {
			primes = append(primes, v)
		}
	}
	fmt.Println("Primes:", primes)

	doExist(42, nums)
	doExist(69, nums)
}
