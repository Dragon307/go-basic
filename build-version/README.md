

## Golang 自动生成版本信息

golang程序在build时自动生成版本信息，使用 ./helloworld -version 可以查看版本和build 时间

## 实现原理
使用链接选项 `-X` 设置一个二进制文件中可以访问的变量

