package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main(){
	var i int
	start := time.Now()
	for i = 0; i < 40; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	times := time.Since(start)
	fmt.Println("times",times)
}