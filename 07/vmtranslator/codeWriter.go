package main

import (
	"os"
	"fmt"
)

type CodeWriter struct{
	file *os.File
}

func NewCodeWriter(file *os.File) *CodeWriter {
	return &CodeWriter{file}
}

func (c *CodeWriter) SetFileName(fileName string) {
}

func (c *CodeWriter) WriteArithmetic(cmd string) {
	switch cmd {
	case "add":
		c.file.WriteString("@SP\n")
		c.file.WriteString("D=M\n")
		c.file.WriteString("AM=D-1\n")
		c.file.WriteString("D=M\n")
		c.file.WriteString("A=A-1\n")
		c.file.WriteString("M=D+M\n")
		c.file.WriteString("\n")
	case "eq":
		c.file.WriteString("@SP\n")
		c.file.WriteString("D=M\n")
		c.file.WriteString("AM=D-1\n")
		c.file.WriteString("D=M\n")
		c.file.WriteString("A=A-1\n")
		c.file.WriteString("D=D-M\n")

		c.file.WriteString("@L0\n")
		c.file.WriteString("D;JEQ\n")
		c.file.WriteString("D=0\n")
		c.file.WriteString("@L1\n")
		c.file.WriteString("0;JMP\n")
		c.file.WriteString("(L0)\n")
		c.file.WriteString("D=-1\n")
		c.file.WriteString("(L1)\n")
		c.file.WriteString("@SP\n")
		c.file.WriteString("A=M\n")
		c.file.WriteString("M=D\n")
		c.file.WriteString("@SP\n")
		c.file.WriteString("M=M+1\n")

		c.file.WriteString("\n")

	}
}

func (c *CodeWriter) WritePushPop(command Command, segment string, index int) {
	switch command {
	case C_PUSH:
		if segment == "constant" {
			c.file.WriteString(fmt.Sprintf("@%d\n", index))
			c.file.WriteString("D=A\n")
			c.file.WriteString("@SP\n")
			c.file.WriteString("A=M\n")
			c.file.WriteString("M=D\n")
			c.file.WriteString("@SP\n")
			c.file.WriteString("M=M+1\n")
			c.file.WriteString("\n")
		}
	case C_POP:
	}
}

func (c *CodeWriter) Close() {
	c.file.Close()
}
