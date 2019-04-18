package main

import (
	"fmt"
	"github.com/cockroachdb/cockroach/pkg/cmd/docgen/extract"
)

func main() {
	b, e := extract.GenerateBNF("/home/robi/Code/go/src/github.com/pingcap/parser/parser.y")
	if e != nil {
		panic(e)
	}
	fmt.Println(string(b))
}
