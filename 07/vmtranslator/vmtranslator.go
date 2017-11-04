package main

import (
	"fmt"
	"path/filepath"
	"os"
	"bufio"
	"strings"
)

func main() {
	dir, inputFile := filepath.Split(os.Args[1])
	fp, err := os.Open(filepath.Join(dir, inputFile))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	p := NewParser(scanner)

	outputFile := strings.Split(inputFile, ".")[0] + ".asm"

	ofp, err := os.OpenFile(filepath.Join(dir, outputFile), os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cw := NewCodeWriter(ofp)
	defer cw.Close()

	for p.HasMoreCommands() {
		p.Advance(p.file.Text())
		if len(p.cmd) == 0 {
			continue
		}
		ct := p.CommandType()
		segment := ""
		if ct != C_RETURN {
			segment = p.Arg1()
		}
		index := 0
		if ct == C_PUSH || ct == C_POP || ct == C_FUNCTION || ct == C_CALL {
			index = p.Arg2()
		}
		if ct == C_PUSH || ct == C_POP {
			cw.WritePushPop(ct, segment, index)
		}
		if ct == C_ARITHMETIC {
			cw.WriteArithmetic(segment)
		}
	}
}
