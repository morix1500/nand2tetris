package main

import (
	"bufio"
	"strings"
	"strconv"
)

type Command int

const (
	C_ARITHMETIC = iota
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
)

type Parser struct {
	file *bufio.Scanner
	cmd string
}

func NewParser(file *bufio.Scanner) *Parser {
	return &Parser {file, ""}
}

func (p *Parser) HasMoreCommands() bool {
	return p.file.Scan()
}

func (p *Parser) Advance(cmd string) {
	p.cmd = ""
	//cmd = strings.Replace(cmd, " ", "", -1)
	comment_index := strings.Index(cmd, "//")
	if comment_index != -1 {
		if comment_index == 0 {
			return
		}
		cmd = strings.Split(cmd, "//")[0]
	}

	if len(cmd) == 0 {
		return
	}
	p.cmd = cmd
	return
}

func (p *Parser) CommandType() Command {
	cmd := strings.Split(p.cmd, " ")[0]

	switch cmd {
	case "push":     return C_PUSH
	case "pop" :     return C_POP
	case "label":    return C_LABEL
	case "goto":     return C_GOTO
	case "if-goto":  return C_IF
	case "function": return C_FUNCTION
	case "return":   return C_RETURN
	case "call":     return C_CALL
	default:         return C_ARITHMETIC
	}
}

func (p *Parser) Arg1() string {
	arr := strings.Split(p.cmd, " ")
	if len(arr) == 1 {
		return p.cmd
	}
	return arr[1]
}

func (p *Parser) Arg2() int {
	arr := strings.Split(p.cmd, " ")
	num, _ := strconv.Atoi(arr[2])
	return num
}
