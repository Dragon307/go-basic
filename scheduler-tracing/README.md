## Scheduler Tracing In Go

GODEBUG 环境变量可以产生运行时的调试信息。你可以请求垃圾回收器和调度器（scheduler）的摘要信息和细节。关键是你不需要额外创建单独的编译
程序就可以实现
schedtrace参数告诉运行时打印一行调度器的摘要信息到标准err输出中，时间间隔可以指定，单位毫秒，如下所示：


> GOMAXPROCS=1 GODEBUG=schedtrace=1000 ./main