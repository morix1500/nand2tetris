// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

    @8192 // 512px * 256px / 16bit
    D=A
    @maxcnt
    M=D
    @nowkbd
    M=0
(LOOP)
    @index
    M=0
    
    @KBD
    D=M
    
    @nowkbd
    M=D
(DISP)
    @nowkbd
    D=M
    
    @BLACK
    D;JGT
(WHITE)
    @index
    D=M
    @SCREEN
    A=A+D
    
    M=0
    @END
    0;JMP
(BLACK)
    @index
    D=M
    @SCREEN
    A=A+D
    M=-1
(END)
    @index
    M=M+1
    
    @maxcnt
    D=M
    @index
    D=D-M
    @DISP
    D;JGT
    
    @LOOP
    0;JMP
