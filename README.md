# go-basic

## GO中的指针

1)  基本数据类型，变量存的就是值，也叫做值类型
2)  获取变量的地址，用&， 比如： var num int, 获取num的地址: &num
3)  指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值
4)  获取指针类型所指向的值，使用：*，比如: var ptr *int, 使用*ptr获取ptr指向的值

* 指针的使用细节

1)  值类型，都有对应的指针类型，形式为 *数据类型，比如 int 的对应的指针就是 *int，float32对应的指针类型就是 *float32
2)  值类型包括： 基本数据类型 int 系列， float 系列，bool，string、数组和结构体 struct

* 值类型和引用类型的说明
1)  值类型: 基本数据类型 int 系列， float系列，bool，string、数组和结构体
2)  引用类型: 指针、slice切片、map、管道chan、interface 等都是引用类型


## 原码、反码、补码


## 位运算符和移位运算符

## 函数的递归调用

## init 函数

## 函数 defer

## forloop