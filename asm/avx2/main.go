package main

import (
	"fmt"
	"golang.org/x/sys/cpu"
)

func main() {
	if cpu.X86.HasAVX2 {
		fmt.Println("cpu sport AVX2")
	}
}
