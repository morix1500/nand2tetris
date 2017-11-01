package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func scanSymbol(p Parser) *SymbolTable {
	address := 0
	st := NewSymbolTable()

	for p.HasMoreCommands() {
		p.Advance(p.file.Text())
		if p.cmd == "" {
			continue
		}
		ct := p.CommandType()
		if ct == A_COMMAND || ct == C_COMMAND {
			address++
		} else {
			sy := p.Symbol(ct)
			if st.Contains(sy) {
				continue
			}
			a := address
			st.AddEntry(sy, a)
		}
	}
	return st
}

func main() {
	dir, input_file := filepath.Split(os.Args[1])
	fp, err := os.Open(filepath.Join(dir, input_file))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	p := NewParser(scanner)
	c := NewCode()

	st := scanSymbol(*p)

	fp.Seek(0, 0)
	scanner = bufio.NewScanner(fp)
	p = NewParser(scanner)

	output := []string{}
	address := 16

	for p.HasMoreCommands() {
		p.Advance(p.file.Text())
		if p.cmd == "" {
			continue
		}
		ct := p.CommandType()
		// binary line
		line := ""

		switch ct {
		case C_COMMAND:
			dest := p.Dest()
			comp := p.Comp()
			jump := p.Jump()

			line = line + "111" + c.Comp(comp) + c.Dest(dest) + c.Jump(jump)
		case A_COMMAND:
			sy := p.Symbol(ct)
			a := 0
			if st.Contains(sy) {
				a = st.GetAddress(sy)
			} else {
				num, err := strconv.Atoi(sy)
				if err == nil {
					a = num
				} else {
					a = address
					st.AddEntry(sy, address)
					address++
				}
			}
			line = line + "0" + fmt.Sprintf("%015b", a)
		default:
			continue
		}
		output = append(output, line)
	}

	output_file := strings.Split(input_file, ".")[0] + ".hack"
	of, err := os.OpenFile(filepath.Join(dir, output_file), os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer of.Close()

	for i := 0; i < len(output); i++ {
		of.WriteString(output[i] + "\n")
	}

	os.Exit(0)
}
