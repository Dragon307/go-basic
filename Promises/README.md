
## Golang 实现Futures和Promises

Future, Promise 或 Delay 是用于并发编程的一种设计模式。它们表示一个对象，这个对象用来作为一次计算结果的代理，
而该结果开始的时候是未知的，因为计算还没有完成。Promise与Future的区别在于，Future是Promise的一个只读的视图，也就是说
Future没有设置任务结果的方法，只能获取任务执行结果或者为Future添加回调函数

