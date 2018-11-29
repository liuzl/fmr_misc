package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/liuzl/fmr"
)

var (
	g = flag.String("g", "test.grammar", "grammar file")
	s = flag.String("s", "number", "start rule")
)

func main() {
	flag.Parse()
	grammar, err := fmr.GrammarFromFile(*g)
	if err != nil {
		log.Fatal(err)
	}
	br := bufio.NewReader(os.Stdin)
	for {
		line, c := br.ReadString('\n')
		if c == io.EOF {
			break
		}
		if c != nil {
			log.Fatal(c)
		}
		line = strings.TrimSpace(line)
		ps, err := grammar.EarleyParseMaxAll(line, *s)
		if err != nil {
			log.Fatal(err)
		}
		for i, p := range ps {
			for _, f := range p.GetFinalStates() {
				trees := p.GetTrees(f)
				fmt.Printf("p%d tree number:%d\n", i, len(trees))
				for _, tree := range trees {
					sem, _ := tree.Semantic()
					nl := tree.NL()
					tag := p.Tag(f)
					fmt.Printf("%s(%s) = %s\n", tag, nl, sem)
				}
			}
		}
		fmt.Println()
	}
}
