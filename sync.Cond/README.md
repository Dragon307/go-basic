## sync.Cond 源码分析
Cond的主要作用就是获取锁之后，wait() 方法会等待一个通知，来进行下一步释放等操作，以此控制锁合适释放，释放频率，适用于在并发环境下goroutine的等待通知

```go
package Cond
import "unsafe"
type Cond struct {
    noCopy noCopy // noCopy 可以嵌套入结构中，在第一次使用后不可复制，使用go vet 作为检测使用
    // 根据需求初始化不同的锁，如*Mutex 和 *RWMutex
    L Locker
    
    notify notifyList // 通知列表，调用Wait() 方法的goroutine会被放入list中，每次唤醒，从这里取出
    checker copyChecker // 复制检查，检查cond实例是否被复制
}

// notifyList 结构体
type notifyList struct {
    wait uint32
    notify uint32
    lock uintptr
    head unsafe.Pointer
    tail unsafe.Pointer
}
```

函数

NewCond
相当于Cond的构造函数，用于初始化 Cond

参数为Locker实例初始化，传参数的时候必须是引用或指针，比如`sync.Mutex{}`或`new(sync.Mutex)`,不然会报错
`cannot use lock (type sync.Mutex) as type sync.Locker in argument to sync.NewCond`
为什么要使用指针，在调用 `c.L.Lock()`和`c.L.Unlock()`的时候，会频繁发生锁的复制，会导致锁的失效，甚至是导致死锁

```go
fun NewCond(l Locker) *Cond {
    return &Cond{L:l}
}
```

Wait

等待自动解锁c.L 和暂停执行调用 goroutine。恢复执行后，等待锁c.L返回之前。于其他系统不同，等待不能返回，除非通过广播或者信号唤醒。
因为c。当等待第一次恢复时，L并没有被锁定，调用者通常不能假定等待返回时条件是正确的。相反，调用者应该在循环中等待

```go
func(c *Cond) Wait() {
    // 检查c是否是被复制的，如果是就panic
    c.checker.check()
    // 将当前goroutine加入等待队列
    t := runtime_notifyListAdd(&c.notify)
    // 解锁
    c.L.Unlock()
    // 等待队列中的所有的goroutine执行等待唤醒操作
    runtime_notifyListWait(&c.notify,t)
    c.L.Lock()
}
```

判断cond是否被复制

```go
type copyChecker uintptr
func (c *copyChecker) check() {
   if uintptr(*c) != uintptr(unsafe.Pointer(c)) && 
    !atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c))) &&
    		uintptr(*c) != uintptr(unsafe.Pointer(c)) {
    		panic("sync.Cond is copied")
}
```

Signal

唤醒等待队列中的一个goroutine，一般都是任意唤醒队列中的一个goroutine，为什么没有选择FIFO模式，因为FiFO模式效率不够高，虽然支持，但是很少使用到

```go
fun (c *Cond) Signal() {
    // 检查c是否是被复制的，如果是就panic
    c.checker.check()
    // 通知等待列表中的一个
    runtime_notifyListNotifyOne(&c.notify)
}
```

Broadcast

唤醒等待队列的所有goroutine

```go
func (c *Cond) Broadcast() {
    // 检查c是否是被复制的，如果是就panic
    c.checker.check()
    // 下发广播给所有等待的goroutine
    runtime_notifyListNotifyAll(&c.notify)
}
```