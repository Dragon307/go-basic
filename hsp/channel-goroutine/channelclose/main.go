package main

import "fmt"

func main()  {

	intChan := make(chan int, 3)
	intChan <- 10
	intChan <- 20

	close(intChan)
	// 通道关闭后就不能够再写入数据，但可以读取数据
	n1 := <- intChan
	fmt.Println("n1=", n1)

	// 遍历管道

	intChan2 := make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}

	// 在遍历时，如果channel 没有关闭，则会出现deadlock的错误
	//在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
