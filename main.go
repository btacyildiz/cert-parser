package main

import (
	"log"
	"os"
)

const (
	notEnoughArguments = "Not enough arguments"
	bundleArgFlag      = "-bundle"
	searchArgFlag      = "-search"
)

type certParserParams struct {
	bundlePath string
	searchText string
}

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal(notEnoughArguments)
	}

	certParserParams := parseArguments(args)

	list, err := parseBundle(certParserParams)
	if err != nil {
		log.Fatalf("unable to parse the bundle: %s", err)
	}

	list.Print()
}

func parseArguments(args []string) certParserParams {
	params := certParserParams{}
	index := 0
	for index < len(args) {
		switch args[index] {
		case bundleArgFlag:
			if index+1 < len(args) {
				index++
				params.bundlePath = args[index]
			}
		case searchArgFlag:
			if index+1 < len(args) {
				index++
				params.searchText = args[index]
			}
		}
		index++
	}
	return params
}
