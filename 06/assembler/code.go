package main

import (
	"strings"
)

type Code struct{}

type CodeInterface interface {
	Dest() string
	Comp() string
	Jump() string
}

func NewCode() *Code {
	return &Code{}
}

func (c *Code) Dest(cmd string) string {
	switch cmd {
	case "M":   return "001"
	case "D":   return "010"
	case "MD":  return "011"
	case "A":   return "100"
	case "AM":  return "101"
	case "AD":  return "110"
	case "AMD": return "111"
	default:    return "000"
	}
}

func (c *Code) Comp(cmd string) string {
	a := "0"
	if strings.Index(cmd, "M") != -1 {
		a = "1"
	}
	comp := ""
	if a == "0" {
		switch cmd {
		case "0":   comp = "101010"
		case "1":   comp = "111111"
		case "-1":  comp = "111010"
		case "D":   comp = "001100"
		case "A":   comp = "110000"
		case "!D":  comp = "001101"
		case "!A":  comp = "110001"
		case "-D":  comp = "001111"
		case "-A":  comp = "110011"
		case "D+1": comp = "011111"
		case "A+1": comp = "110111"
		case "D-1": comp = "001110"
		case "A-1": comp = "110010"
		case "D+A": comp = "000010"
		case "D-A": comp = "010011"
		case "A-D": comp = "000111"
		case "D&A": comp = "000000"
		case "D|A": comp = "010101"
		}
	} else {
		switch cmd {
		case "M":   comp = "110000"
		case "!M":  comp = "110001"
		case "-M":  comp = "110011"
		case "M+1": comp = "110111"
		case "M-1": comp = "110010"
		case "D+M": comp = "000010"
		case "D-M": comp = "010011"
		case "M-D": comp = "000111"
		case "D&M": comp = "000000"
		case "D|M": comp = "010101"
		}
	}

	return a + comp
}

func (c *Code) Jump(cmd string) string {
	switch cmd {
	case "JGT": return "001"
	case "JEG": return "010"
	case "JGE": return "011"
	case "JLT": return "100"
	case "JNE": return "101"
	case "JLE": return "110"
	case "JMP": return "111"
	default:    return "000"
	}
}
