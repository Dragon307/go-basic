## sync.Cond 源码分析
Cond的主要作用就是获取锁之后，wait() 方法会等待一个通知，来进行下一步释放等操作，以此控制锁合适释放，释放频率，适用于在并发环境下goroutine的等待通知

```go
package Cond
type Cond struct {
    noCopy noCopy // noCopy 可以嵌套入结构中，在第一次使用后不可复制，使用go vet 作为检测使用
    // 根据需求初始化不同的锁，如*Mutex 和 *RWMutex
    L Locker
    
    notify notifyList // 通知列表，调用Wait() 方法的goroutine会被放入list中，每次唤醒，从这里取出
    checker copyChecker // 复制检查，检查cond实例是否被复制
}
```
