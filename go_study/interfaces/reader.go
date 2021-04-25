package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(input []byte) (int, error) {
	capInput := len(input)
	for i := 0; i < capInput; i++ {
		input[i] = 'A'
	}
	return capInput, nil
}

func main() {
	reader.Validate(MyReader{})
}
