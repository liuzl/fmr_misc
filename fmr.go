package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/liuzl/fmr"
)

var (
	g = flag.String("g", "test.grammar", "grammar file")
)

func main() {
	flag.Parse()
	grammar, err := fmr.GrammarFromFile(*g)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(grammar)
}
