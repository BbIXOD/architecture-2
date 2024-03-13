package main

import (
	"flag"
	"io"
	"os"
	"strings"

	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFileName   = flag.String("f", "", "Input file")
	outputFileName  = flag.String("o", "", "Output file")

	inputFile  io.Reader
	outputFile io.Writer
)

func main() {
	flag.Parse()

	if *inputExpression == "" && *inputFileName == "" {
		panic("no input")
	} else if *inputExpression != "" && *inputFileName != "" {
		panic("too many inputs")
	}

	if *inputFileName != "" {
		inputFile = tryActionWithFile(os.Open, *inputFileName)
	} else if *inputExpression != "" {
		inputFile = strings.NewReader(*inputExpression)
	}

	if *outputFileName != "" {
		outputFile = tryActionWithFile(os.Create, *outputFileName)
	} else {
		outputFile = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  inputFile,
		Output: outputFile,
	}
	err := handler.Compute()

	if err != nil {
		panic(err)
	}
}

func tryActionWithFile(action func(string) (*os.File, error), fileName string) *os.File {
	f, err := action(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return f
}
