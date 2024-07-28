package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CalcFromFile(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		return
	}

	calculator := NewCalculatorVisitor()
	calculator.performExpression(string(content))
}

func CalcFromParam(params []string) {
	expression := strings.Join(params, " ")
	expression += "\n"

	calculator := NewCalculatorVisitor()
	calculator.performExpression(expression)
}

func CalcInteractive() {
	calculator := NewCalculatorVisitor()
	fmt.Println("Entering interactive mode. Type 'exit' to quit.")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "exit" {
			break
		}
		line += "\n"
		calculator.performExpression(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
