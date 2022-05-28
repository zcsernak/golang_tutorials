package main

import "fmt"

func main() {
	// when implementing interfaces,
	// if any of the methods require a pointer receiver
	// then the interface has to be implemented with a pointer:
	var wc WriterCloser = &myWriterCloser{}
	fmt.Println(wc)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type myWriterCloser struct{}

// pointer receiver
func (bwc *myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

// value receiver
func (bwc myWriterCloser) Close() error {
	return nil
}
