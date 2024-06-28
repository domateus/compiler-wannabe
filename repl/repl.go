package repl

import (
	"bufio"
	"fmt"
	"interpreter/lexer"
	"interpreter/parser"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	loop(in, out, scanner)
}

func loop(in io.Reader, out io.Writer, scanner *bufio.Scanner) {
	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recuperaçao de pânico. Erro:\t", r)
				fmt.Println("Programa nao aceito")
				clear(p.Errors())
				loop(in, out, scanner)
			}
		}()

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			for _, msg := range p.Errors() {

				io.WriteString(out, "\n")
				io.WriteString(out, "\n")
				io.WriteString(out, "\t"+msg+"\n")
				io.WriteString(out, "\n")
				io.WriteString(out, "\n")
			}
		}

		if len(p.Errors()) > 0 {
			msg := fmt.Sprintf("O progrma tinha: %d erros\n", len(p.Errors()))
			io.WriteString(out, msg)
			clear(p.Errors())
		} else {

			io.WriteString(out, program.String())
			io.WriteString(out, "\n")
		}
	}
}
