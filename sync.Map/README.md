
## sync.Map 分析

go 普通的map是不支持并发，换而言之，不是线程(goroutine)安全的。博主是从 golang 1.4 开始使用的，那时候map的并发读是没有支持，
但是并发写会出现脏数据。

所以需要支持对map的并发读写时候，有两种处理方法

1，使用第三方类库 [concurrent-map](https://github.com/orcaman/concurrent-map)
2, map加上 `sync.RWMutex` 来保障线程(goroutine) 安全的

## sync.Map

```go
type Map struct {
   mu Mutex // 互斥锁，用于锁定 dirty map
   read atomic.Value // 优先读map，支持原子操作，注释中有readOnly 不是说 read 是只读，而是它的结构体。read 实际上有写的操作
   dirty map[interface{}]*entry // dirty 是一个当前最新的map，允许读写
    misses int // 主要记录read 读取不到数据加锁读取 read map以及dirty map的次数，当misses等于dirty的长度时，会将dirty复制到read
}
```

readOnly

readOnly 主要用于存储，通过原子操作存储在Map.read中元素。

```go
type readOnly struct {
    m map[interface{}]*entry
    amended bool // 如果数据在dirty 中但没有在read中，该值为true，作为修改标识
}
```

entry

```go
type entry struct {
    // nil: 表示为被删除，调用Delete() 可以将 read map中的元素置为nil
    // expunged: 也是表示被删除，但是该键只在read 而没有在dirty中，这种情况出现将 read 复制到 dirty 中，即复制的过程会先将nil 标记为expunged, 然后不将其复制到dirty
    // 其他：表示存着真正的数据
    p unsafe.Pointer
}
```