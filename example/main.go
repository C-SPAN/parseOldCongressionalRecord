package main

import (
	"fmt"
	"log"
	"os"

	pocr "github.com/C-SPAN/parseOldCongressionalRecord"
)

func main() {
	var txt string

	args := os.Args

	if len(args) < 3 {
		log.Fatal("missing required arguments: parseOldCongressionalRecord <inFile> <outFile>")
	}

	log.Printf("reading %s to %s\n", args[1], args[2])

	xmlData, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatal(fmt.Errorf("error reading file: %w", err))
	}

	txt, err = pocr.ParseXML(string(xmlData))
	if err != nil {
		log.Fatal(fmt.Errorf("error parsing xml: %w", err))
	}

	err = os.WriteFile(args[2], []byte(txt), 0644)
	if err != nil {
		log.Fatal(fmt.Errorf("error writing to file: %w", err))
	}

	log.Println("done")
}
