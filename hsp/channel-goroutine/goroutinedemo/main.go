package main

import (
	"fmt"
	"sync"
)

var (
	myMap = make(map[int]int, 10)

	lock sync.Mutex
)

func test(n int)  {
	res := 1
	for i:= 1; i<= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {

	// 我们这里开启多个协程完成这个任务[200个]
	for i := 1; i <= 20; i++ {
		go test(i)
	}


	//休眠10秒钟【第二个问题 】
	//time.Sleep(time.Second * 5)

	//这里我们输出结果,变量这个结果
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()

}