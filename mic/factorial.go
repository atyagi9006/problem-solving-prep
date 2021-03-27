package main

import "fmt"

func main() {
	fmt.Println("FActR 15", factR(15))
	fmt.Println("FAct 15", fact(15))
}

func factR(n int) int {
	if n == 1 {
		return 1
	}
	return n * factR(n-1)
}

func fact(n int) int {
	if n == 1 {
		return 1
	}
	fact := 1
	for i := n; i > 0; i-- {
		fact = fact * i
	}
	return fact
}


