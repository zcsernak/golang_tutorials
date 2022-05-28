package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var wc WriterCloser = NewBufferedWriterCloser()
	n, err := wc.Write([]byte("Hello YouTube listeners, this is a test"))
	if err == nil {
		/*defer */ wc.Close()
		fmt.Println("Wrote", n, "bytes")
	} else {
		panic(err)
	}

	fmt.Println()
	typeConversionExample()
	fmt.Println()
	emptyInterfaceExample()
	fmt.Println()
	typeSwitchExample()
}

func typeConversionExample() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello YouTube listeners, this is a test"))
	wc.Close()

	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	// bwc := wc.(io.Reader) // panic: interface conversion: *main.BufferedWriterCloser is not io.Reader: missing method Read

	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}
}

func emptyInterfaceExample() {
	// empty interface: interface{}
	// everything can be casted to an empty interface
	var myObj interface{} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("Hello YouTube listeners, this is a test"))
		wc.Close()
	}
	r, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}
}

func typeSwitchExample() {
	var i interface{} = "0"
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// interface composed of 2 other interfaces
type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)

	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data)) // print remaining bytes
		if err != nil {
			return err
		}
	}
	return nil
}

// Constructor method for BufferedWriterCloser
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
