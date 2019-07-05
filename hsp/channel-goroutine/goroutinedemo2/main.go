package main

import "fmt"

//var myMap = make(map[int]int, 10)

var intChan = make(chan int, 10)

func test(n int)  {
	res := 1
	for i:= 1; i<= n; i++ {
		res *= i
	}
	intChan <- res
}

func main()  {
	for i := 1; i <= 20; i++ {
		go test(i)
	}
forloop:
	for {
		select {
		case n := <-intChan:
			fmt.Println(n)
		default:
			fmt.Println("no more")
			break forloop

		}
	}

}