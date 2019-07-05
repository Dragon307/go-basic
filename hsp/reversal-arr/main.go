package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	var intArr [6]int
	len := len(intArr)

	rand.Seed(time.Now().UnixNano())
	//Seed uses the provided seed value to initialize the generator to a deterministic state.
	// Seed should not be called concurrently with any other Rand method.
	for i := 0; i< len; i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Println("反转前==", intArr)

	temp := 0
	for i := 0; i< len/2; i++ {
		temp = intArr[len-1-i]
		intArr[len-1-i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println("反转后==", intArr)
}