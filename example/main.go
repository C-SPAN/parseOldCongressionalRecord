package main

import (
	"log"
	"os"

	pocr "github.com/C-SPAN/parseOldCongressionalRecord"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		log.Fatal("missing required arguments: parseOldCongressionalRecord <inFile> <outFile>")
	}

	txt, err := pocr.ParseXML(args[1])
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(args[2], []byte(txt), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
