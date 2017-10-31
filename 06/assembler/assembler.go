package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"path/filepath"
)

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

	output := []string{}

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
			v, err := strconv.Atoi(sy)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			line = line + "0" + fmt.Sprintf("%015b", v)
		case L_COMMAND:
			sy := p.Symbol(ct)
			v, err := strconv.Atoi(sy)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			line = line + "111" + fmt.Sprintf("%013b", v)
		}
		output = append(output,line)
	}

	output_file := strings.Split(input_file, ".")[0] + ".hack"
	of, err := os.OpenFile(filepath.Join(dir, output_file), os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer of.Close()

	for i := 0 ; i < len(output) ; i++ {
		of.WriteString(output[i] +  "\n")
	}

	os.Exit(0)
}
