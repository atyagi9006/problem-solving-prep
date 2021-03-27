package main

import (
	"fmt"
	"sync"
)

func executeParallel(ch chan<- int, functions ...func() int) {
	var wg sync.WaitGroup
	fLen:=len(functions)
	fmt.Println()
	wg.Add(fLen)
	for i:=0;i<fLen-1;i++{
		go func() {
			res:= functions[i]()
			fmt.Printf("Result L1: %v\n", functions[i]())
			ch <- res
			wg.Done()
		}()
	}
		
		/* go func() {
			res:= functions[1]()
			fmt.Printf("Result L1: %v\n", functions[1]())
			ch <- res
			wg.Done()
		}() */
	//}
	wg.Wait()
	close(ch)
}

func exampleFunction(counter int) int {
	sum := 0
	for i := 0; i < counter; i++ {
		sum += 1
	}
	return sum
}

func main() {
	expensiveFunction := func() int { return exampleFunction(200000000) }
	cheapFunction := func() int { return exampleFunction(10000000) }

	ch := make(chan int)

	go executeParallel(ch, expensiveFunction, cheapFunction)

	for result := range ch {
		fmt.Printf("Result: %d\n", result)
	}
}
