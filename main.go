package main

import (
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error: ", err) // Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}
	defer file.Close()
}
