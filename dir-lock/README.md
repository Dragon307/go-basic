
## golang 的文件锁操作

我们在使用golang开发程序的时候，经常会出现多个goroutine操作同一个文件（或目录）的时候，如果不加锁，很容易导致文件中的数据混乱，于是，Flock应运而生

Flock 是对于整个文件的建议性锁（不强求 goroutine 遵守），如果一个 goroutine 在文件上获取了锁，那么其他goroutine是可以知道的。默认情况下，当一个
goroutine 将文件锁住，另外一个 goroutine 可以直接操作被锁住的文件，原因在于Flock只是用于检测文件是否被加锁，针对文件已经被加锁，另一个goroutine写入
数据的情况，内核不会阻止这个 goroutine 的写入操作，也就是建议性锁的内核处理策略

## 函数

```go
import "syscall"
func Flock(fd int, how int)(err error)  
```

Flock 位于 syscall 包中，fd参数指代文件描述符，how参数指代锁的操作类型

how 主要的参数类型：

* LOCK_SH, 共享锁，多个进程可以使用同一把锁，常被用作读共享锁；
* LOCK_EX, 排它锁，同时只允许一个进程使用，常被用作写锁；
* LOCK_NB, 遇到锁的表现，当采用排它锁的时候，默认goroutine会被阻塞等待被释放，采用 LOCK_NB 参数，可以让 goroutine 返回 Error；
* LOCK_UN, 释放锁；