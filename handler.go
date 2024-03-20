package lab2

import (
	"fmt"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {

	input, err := io.ReadAll(ch.Input)
	if err != nil {
		return err
	}

	result, err := СalculatePrefix(string(input))
	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(fmt.Sprintf("%d", result)))
	if err != nil {
		return err
	}
	return nil
}
