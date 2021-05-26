package main

import (
	"fmt"
	"io"
	"os"
)

// Read file and return the contents.

func main() {
	fileName := os.Args[1]

	f, error := os.Open(fileName)

	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)

}
