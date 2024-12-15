package main

import (
	"fmt"
	"io"
	"os"
)

type writer struct{}

func main() {
	// fmt.Println(os.Args[1])

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error is : ", err)
		os.Exit(1)
	}

	w := writer{}

	// io.Copy(os.Stdout, f)
	io.Copy(w, f)

}

func (writer) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Bytes processed is : ", bs)
	return len(bs), nil
}
