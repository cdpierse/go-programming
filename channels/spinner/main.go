package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 50
	fibN := fib2(n)
	fmt.Printf("\r Fibonnaci(%d) = %d\n", n, fibN)

}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)

}

func fib2(x int) int {
	fib := []int{0,1}
	// fib = append(fib,1)
	// fib = append(fib,2)
	

	for i:=0; i < x-1; i++ {
		fmt.Println(fib)
		fib[0], fib[1] = fib[1],fib[0] + fib[1]
			
	}
	if x == 0{
		return 0
	}

	return fib[1]
	
}

