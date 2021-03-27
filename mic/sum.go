package main

import "fmt"

func main() {
	fmt.Println("SumR N", sumR(15))
	fmt.Println("Sum N", sum(15))
}

func sumR(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return n + sumR(n-1)
}

func sum(n int) int {
	sum := 0
	if n == 0 {
		return sum
	}

	for i := 1; i < n+1; i++ {
		sum = sum + i
	}

	return sum
}
