package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func main() {
	for i:=0; i< 40; i++ {
		go func(x int) {
			cond.L.Lock() // 获取锁
			defer cond.L.Unlock() // 释放锁
			cond.Wait()// 等待通知，阻塞当前 goroutine
			fmt.Println(x)
			time.Sleep(time.Millisecond * 300)
		}(i)
	}

	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal() // 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	cond.Signal() // 3 秒之后下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	cond.Broadcast()
	fmt.Println("Broadcast...") // 3 秒之后 下发广播给所有等待的 goroutine

	time.Sleep(time.Second * 20)
}