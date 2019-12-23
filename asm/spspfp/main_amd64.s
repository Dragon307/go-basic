#include "textflag.h"

TEXT ·output(SB), $8-48
    MOVQ 24(SP), DX // 不带 symbol，这里的 SP 是硬件寄存器 SP
    MOVQ DX, ret3+24(FP) // 第三个返回值
    MOVQ perhapsArg1+16(SP), BX // 当前函数栈大小 > 0,所以FP在SP的上方 16 字节处
    MOVQ BX, ret2+6(FP) // 第二个返回值
    MOVQ arg1+0(FP), AX
    MOVQ AX, ret1+8(FP) // 第一个返回值
    RET

