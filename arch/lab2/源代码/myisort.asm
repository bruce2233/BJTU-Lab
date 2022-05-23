.data
array: .word  0x65321010,0x11111111,0x00000000,0x22222222
       .word  0x44444444,0x55555555,0x77777777,0x33333333

len: .word 16

.text
        addi t0,zero,4
        addi x20,zero,2
        la x18,array
        la s6,len
        lw t1,0(s6)
        sll t1,t1,x2
for:    slt t2,t0,t1
        beqz t2,out
        add t3,zero,t0
        addi x18,x18,4
        lw t4,0(x18)
        addi x19,x18,0
loop:   slt t2,zero,t3
        beqz t2,over
        
        addi t5,t3,-4
        lw t6,-4(x19)
        
        slt t2,t6,t4
        beqz t2,over
        sw t6,0(x19)
        addi x19,x19,-4
        add t3,zero,t5
        
        j loop

over:   sw t4,0(x19)
        addi t0,t0,4
        j for
out:    ecall