

[x86 指令集](https://github.com/golang/arch/blob/master/x86/x86.csv)

[go 定义的指令集别名](https://golang.org/src/cmd/internal/obj/x86/anames.go)



## Go Assembly 

https://xargin.com/plan9-assembly/

[Go Assembly 示例](https://colobu.com/goasm/)

[A Quick Guide to Go's Assembler](https://golang.org/doc/asm)

[汇编入门](https://chai2010.cn/advanced-go-programming-book/ch3-asm/readme.html)

## [plan9 汇编简介](https://zhuanlan.zhihu.com/p/56750445)

主要包含: 文件命名、指令集、寄存器、函数声明、全局变量声明、跳转、栈分布、调用栈、编译/反编译工具

### 文件命名

由于不同的平台架构支持的汇编指令集不一样，需要针对不同的架构写不同的汇编实现。
通常文件命名格式： 功能名_arch.s
比如：indexbyte_386.s, indexbyte_arm64.s, indexbyte_s390x.s
使用go build编译的时候，会自动根据当前arch平台使用对应的arch文件(或者使用 + build tag)

### 指令集
MAC 查看指令集信息 machdep.cpu.features:

> sysctl machdep.cpu

### 寄存器

有4个核心的伪寄存器，这4个寄存器是编译器用来维护上下文、特殊标识等作用的:

FP(Frame pointer): arguments and locals
PC(Program counter): jumps and branches
SB(Static base pointer): global symbols
SP(Stack pointer): top of stack

所有用户空间的数据都可以通过`FP/SP`（局部数据、输入参数、返回值）和 `SB`(全局数据)访问。通常情况下，
不会对SB/FP寄存器进行运算操作，通常情况以会以 SB/FP/SP 作为基准地址，进行偏移解引用等操作。

其中

1. SP 有伪SP和硬件SP的区分：

伪SP： 本地变量最高起始地址

硬件SP: 函数栈真实栈顶地址

他们的关系是：
* 如果没有本地变量: 伪SP = 硬件SP + 8
* 如果有本地变量: 伪SP = 硬件SP + 16 + 本地变量空间大小

2.  FP伪寄存器

FP伪寄存器： 用来标识函数参数、返回值和伪SP寄存器的关系是：

* 如果有本地变量或者栈调用存严格split在关系（无NOSPLIT），伪FP = 伪SP + 16 
* 否则，伪FP = 伪SP + 8
* FP是访问入参、出参的基址，一般用正向偏移来寻址，SP是访问本地变量的起始基址，一般用负向偏移来寻址
* 修改硬件SP，会引起SP、FP同步变化

> SUBQ $16, SP  // 这里 golang解引用，伪SP/FP 都会-16

3.  参数/本地变量访问

通过`symbol+/-offset(FP/SP)` 的方式进行使用
例如arg0+0(FP) 表示第函数第一个参数起始的位置，arg1+8(FP)表示函数参数偏移8byte的另一个参数。
arg0/arg1用于助记，但是必须存在，否则无法编译通过（golang会识别并做处理）。
其中对于SP来说，还有一种访问方式:
+/-offset(FP)
这里SP前面没有symbol修饰，代表这硬件SP

4.  PC寄存器

实际上就是在体系结构的只是中常见的pc寄存器，在x86平台下对应ip寄存器，amd64上则是rip。
除了个别跳转之外，手写代码与PC寄存器打交道的情况较少

5.  SB寄存器

SP是栈指针寄存器，指向当前函数栈的栈顶，通过symbol+offset(SP)的方式使用。offset 的合法取值是[-framesize, 0), 注意这个是左闭右开区间。
假如局部变量都是 8 字节，那么第一个局部变量就可以用localvar0-8(SP)来表示

6.  BP寄存器

还有BP寄存器，是表示已给调用栈的起始栈底（栈的方向从大到小，SP表示栈顶）； 一般用的不多，如果需要做手动维护调用栈

7:  通用寄存器

在plan9汇编里还可以直接使用的amd64的通用寄存器，应用代码层面会用到的通用寄存器主要是：rax，rbx，rcx，rdx，rdi，rsi，r8~r15 这14个寄存器。
plan9中使用寄存器不需要带r或e的前缀，例如rax，只要写AX即可：

>   MOVQ $101, AX

代码示例

```go
func Add(a, b int) (c int) {
    sum := 0
    return a + b + sum
}
```
各个变量通过寄存器解引用如下: (伪FP = 伪SP+16 = 硬件SP+24)

* a： a+0(SP) 或者 a+16(SP)
* b:  b+8(SP) || a+ 24(SP)
* c:  c+16(SP) || a + 32(SP)
* sum: sum-8(SP) || a - 24(FP)

### 函数声明

```go
package main

import "math"

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
```

```text
"".Sqrt STEXT nosplit size=17 args=0x10 locals=0x0
        0x0000 00000 (hello.go:5)       TEXT    "".Sqrt(SB), NOSPLIT|ABIInternal, $0-16
        0x0000 00000 (hello.go:5)       FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (hello.go:5)       FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (hello.go:5)       FUNCDATA        $2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (hello.go:6)       PCDATA  $0, $0
        0x0000 00000 (hello.go:6)       PCDATA  $1, $0
        0x0000 00000 (hello.go:6)       MOVSD   "".x+8(SP), X0
        0x0006 00006 (hello.go:6)       SQRTSD  X0, X0
        0x000a 00010 (hello.go:6)       MOVSD   X0, "".~r1+16(SP)
        0x0010 00016 (hello.go:6)       RET
```
此处声明了一个函数sqrt，函数的声明以 TEXT 标识开头，以 {package}·{function} 为函数名。 如何函数属于本package时，通常可以不写{package}，
只留·{function}即可。· 在mac上可以用shift+option+9 打出。$0表示该函数栈大小为0byte，计算栈大小时，需要考虑局部变量和本函数内调用其他函数时，
需要传参的空间，不含函数返回地址和CALLER BP。 $16表示该函数入参和返回值一共有16byte。当有NOSPLIT标识时，可以不写输入参数、返回值占用的大小
（这时候会强行插入CALLER BP）。

### 全局变量声明

全局变量的数据部分采用DATA symbol+offset(SB)/width, value格式进行声明。<>表示该变量只在该文件内全局可见。

```text
DATA divtab<>+0x00(SB)/4, $0xf4f8fcff  // divtab的前4个byte为0xf4f8fcff
DATA divtab<>+0x04(SB)/4, $0xe6eaedf0  // divtab的4-7个byte为0xe6eaedf0
...
DATA divtab<>+0x3c(SB)/4, $0x81828384  // divtab的最后4个byte为0x81828384
GLOBL divtab<>(SB), RODATA, $64        // 全局变量名声明，以及数据所在的段"RODATA"，数据的长度64byte
```

### 跳转

跳转分为section跳转或者函数调用跳转

* section 跳转

类似JNE，JBE，JE，JGE等；其中sp/bp不会变化；栈空间不变，不存在参数传递需求

* 函数调用跳转

JMQ sp/bp 不会变化；栈空间不变。通常需要调用者和被调用者协商好使用那些寄存传递参数，调用者将参数写入这些寄存器

CALL 栈空间会发生响应的变化，传递参数时，我们需要输入参数、返回值按之前将的栈布局安排在调用者的栈顶(低地址段)，然后再调用CALL
命令来调用其函数，调用CALL命令后，SP寄存器会下移一个WORD(x86_64上是8byte)，然后进入新函数的栈空间运行

### 编译/ 反编译工具

编译

>   go build -gcflags="-S"
> go tool compile -S hello.go
> go tool compile -N -S hello.go // 禁止优化

反编译

> go tool objdump <binary>


## 工具

使用命令查看Go语言程序对应的伪汇编代码

> go tool compile -S main.go