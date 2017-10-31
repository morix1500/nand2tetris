package main

import(
	"bufio"
	"strings"
)
type Command int

const (
	A_COMMAND Command = iota
	C_COMMAND
	L_COMMAND
)

type ParserInterface interface {
	HasMoreCommands() bool
	Advance()
	CommandType() Command
	Symbol() string
	Dest() string
	Comp() string
	Jump() string
}

type Parser struct {
	file  *bufio.Scanner
	cmd string
}

func NewParser(scanner *bufio.Scanner) *Parser {
	return &Parser{scanner, ""}
}

func (p *Parser) HasMoreCommands() bool {
	return p.file.Scan()
}

func (p *Parser) Advance(cmd string) {
	p.cmd = ""
	cmd = strings.Replace(cmd, " ", "", -1)
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
	if strings.HasPrefix(p.cmd, "@") {
		return A_COMMAND
	}
	if strings.HasPrefix(p.cmd, "(") {
		return L_COMMAND
	}
	return C_COMMAND
}

func (p *Parser) Symbol(ct Command) string {
	switch ct {
	case A_COMMAND:
		return p.cmd[1:]
	case L_COMMAND:
		return p.cmd[1:len(p.cmd)-2]
	}

	return ""
}

func (p *Parser) Dest() string {
	if strings.Index(p.cmd, "=") != -1 {
		return strings.Split(p.cmd, "=")[0]
	}
	return ""
}

func (p *Parser) Comp() string {
	str := p.cmd
	if strings.Index(str, "=") != -1 {
		str = strings.Split(str, "=")[1]
	}
	if strings.Index(p.cmd, ";") != -1 {
		str = strings.Split(str, ";")[0]
	}
	return str
}

func (p *Parser) Jump() string {
	if strings.Index(p.cmd, ";") != -1 {
		arr := strings.Split(p.cmd, ";")
		return arr[len(arr) - 1]
	}
	return ""
}

