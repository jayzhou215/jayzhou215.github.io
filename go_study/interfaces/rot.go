package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(input []byte) (int, error) {
	// 原始input仅是个未赋值的byte array
	readLen, err := r.r.Read(input)
	if err != nil {
		return readLen, err
	}
	for i := 0; i < readLen; i++ {
		rot13(input, i)
		//input[i] = rot13_2(input[i])
	}
	return readLen, err
}

func rot13(input []byte, i int) {
	inputByte := input[i]
	// 这里是基于比较
	if inputByte >= 'A' && inputByte <= 'M' {
		input[i] = inputByte + 13
	} else if inputByte >= 'N' && inputByte <= 'Z' {
		input[i] = inputByte - 13
	} else if inputByte >= 'a' && inputByte <= 'm' {
		input[i] = inputByte + 13
	} else if inputByte > 'm' && inputByte <= 'z' {
		input[i] = inputByte - 13
	} else {
		input[i] = inputByte
	}
}

// https://gist.github.com/flc/6439105 这个示例的代码里面是基于模运算

var ascii_uppercase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var ascii_lowercase = []byte("abcdefghijklmnopqrstuvwxyz")
var ascii_uppercase_len = len(ascii_uppercase)
var ascii_lowercase_len = len(ascii_lowercase)

func rot13_2(b byte) byte {
	pos := bytes.IndexByte(ascii_uppercase, b)
	if pos != -1 {
		return ascii_uppercase[(pos+13)%ascii_uppercase_len]
	}
	pos = bytes.IndexByte(ascii_lowercase, b)
	if pos != -1 {
		return ascii_lowercase[(pos+13)%ascii_lowercase_len]
	}
	return b
}

func main() {
	//bytes := []byte{'a', 'b', 'z', 'A', 'Z'}
	//fmt.Printf("%v, %d, %d, %d, %d", string(bytes), bytes[0], bytes[2], bytes[3], bytes[4])
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
