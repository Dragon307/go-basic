TEXT ·Fab(SB),$0-16
    MOVQ n+0(FP), AX
    CMPQ AX, $1
    JBE end
    SUBQ $16, SP
    MOVQ AX, 0(SP)
    DECQ 0(SP)
    CALL ·Fab(SB)
    MOVQ 8(SP), AX
    MOVQ AX, 40(SP)
    DECQ 0(SP)
    CALL ·Fab(SB)
    MOVQ 8(SP), AX
    ADDQ AX, 40(SP)
    ADDQ $16, SP
    RET
end:
    MOVQ $1, ret+8(FP)
    RET