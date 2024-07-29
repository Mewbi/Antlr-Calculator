package main

import (
	"flag"
	"fmt"

	"calculator/calculator"
)

func main() {
	// Define flags
	fileFlag := flag.String("file", "", "Read input from the specified file")
	helpFlag := flag.Bool("help", false, "Help menu")
	flag.Parse()

	args := flag.Args()

	if *helpFlag {
		showHelp()
		return
	}

	// If file is passed
	if len(*fileFlag) > 0 {
		calculator.CalcFromFile(*fileFlag)
		return
	}

	// If an expression is passed as an argument
	if len(args) > 0 {
		calculator.CalcFromParam(args)
		return
	}

	// Interactive prompt mode
	calculator.CalcInteractive()
}

func showHelp() {
	fmt.Println("Usage:")
	fmt.Println("  ./calc [expression]        Evaluate the provided expression")
	fmt.Println("  ./calc -file [file]        Read and print the content of the specified file")
	fmt.Println("  ./calc                     Run in interactive prompt mode")
	fmt.Println("  ./calc -help               Show this help message")
}
