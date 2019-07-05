package main

import "fmt"

func main() {

	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	intChan<- 10
	num := 211
	intChan<- num
	intChan<- 50

	<- intChan
	intChan<- 98

	//4. 看看管道的长度和cap(容量)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 3, 3

	//5. 从管道中读取数据

	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan))  // 2, 3

	//6. 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	num3 := <-intChan
	num4 := <-intChan

	//num5 := <-intChan

	fmt.Println("num3=", num3, "num4=", num4/*, "num5=", num5*/)

}
