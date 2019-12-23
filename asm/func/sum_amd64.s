
TEXT  ·Sum+0(SB),$0
    MOVQ a+0(FP), BX
    MOVQ b+8(FP), BP
    ADDQ BP, BX
    MOVQ BX, return+16(FP)
    RET
