package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Reading file: %s \n", file.Name())

	io.Copy(os.Stdout, file)
}
