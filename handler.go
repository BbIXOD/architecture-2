package lab2

import (
	"bufio"
	"fmt"
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

	result, err := СalculatePrefix(input)
	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(fmt.Sprintf("%d", result)))
	if err != nil {
		return err
	}
	return nil
}
