package main

import (
	"fmt"
)

// write Data
func writeData(intChan chan int)  {
	for i := 1; i< 50; i++ {
		intChan<- i
		fmt.Println("writeData",i)
		// time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool)  {

	for {
		v, ok := <- intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	// readData 读取完数据后
	exitChan<-true
	close(exitChan)
}

func main()  {
	// 创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
