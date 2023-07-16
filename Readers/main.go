package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(p []byte) (int, error) {

	return copy(p, []byte{'A'}), nil
}

func main() {
	reader.Validate(MyReader{})
}
