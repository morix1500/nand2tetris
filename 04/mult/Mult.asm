// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

    @i
    M=0

    @sum
    M=0
(LOOP)
    @i
    D=M

    @R1
    D=D-M

    @END
    D;JGE // if (i - R1) <= 0

    @R0
    D=M

    @sum
    M=M+D // sum += R0

    @i
    M=M+1 // i++
 
    @LOOP
    0;JMP // to the top
(END)
    @sum
    D=M

    @R2
    M=D
