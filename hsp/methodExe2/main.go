package main

import "fmt"

type MethodUtils struct {

}

func (mu MethodUtils) Print()  {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) Print2(m int, n int)  {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) area(len float64, width float64) float64 {
	return  len * width
}

func main()  {
	var mu MethodUtils
	mu.Print()
	mu.Print2(5, 20)

	areaRes := mu.area(5.1, 2.0)
	println(areaRes)
}