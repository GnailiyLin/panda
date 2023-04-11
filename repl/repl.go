package repl

import (
	"bufio"
	"fmt"
	"io"
	"panda/lexer"
	"panda/parser"
)

const Prompt = ">>"

const Panda = `
   ________    ________    ________     ______     ________ 
  /        \  /        \  /    /   \  _/      \\  /        \
 /         / /         / /         / /        // /         /
//      __/ /         / /         / /         / /         / 
\\_____/    \___/____/  \__/_____/  \________/  \___/____/  

`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	fmt.Fprintf(out, Panda)

	for {
		fmt.Fprintf(out, Prompt)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "  parser errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
