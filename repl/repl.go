package repl

import (
	"bufio"
	"fmt"
	"io"
	"panda/compiler"
	"panda/lexer"
	"panda/parser"
	"panda/vm"
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

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)

			continue
		}

		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)

			continue
		}

		stackTop := machine.LastPoppedStackElem()
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "  parser errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
