package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, err error) {
	z := 0
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
		z++
	}

	return z, nil
}

func main() {
	reader.Validate(MyReader{})
}
