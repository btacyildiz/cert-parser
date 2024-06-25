package main

import (
	"log"
	"os"
)

const (
	notEnoughArguments = "Not enough arguments"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal(notEnoughArguments)
	}

	list, err := parseBundle(args[0])
	if err != nil {
		log.Fatalf("unable to parse the bundle: %s", err)
	}

	list.Print()
}
