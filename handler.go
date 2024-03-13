package lab2

import (
	"bufio"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	reader := bufio.NewReader(ch.Input)

	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	result, err := PrefixToPostfix(input)
	if err != nil {
		return err
	}

	ch.Output.Write([]byte(result))
	return nil
}
