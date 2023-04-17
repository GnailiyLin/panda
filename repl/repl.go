package repl

import (
	"bufio"
	"fmt"
	"io"
	"panda/evaluator"
	"panda/lexer"
	"panda/object"
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
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

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

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "  parser errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
