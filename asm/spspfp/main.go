package main

import "fmt"

func output(int) (int, int, int) // 汇编函数声明

func main() {
	a, b, c := output(987654321)
	fmt.Println(a, b, c)
}

// 证明伪SP，伪FP和硬件SP的位置关系