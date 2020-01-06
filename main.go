package main

import (
	"flag"
	"log"

	"net.fujlog/go-wasm/parser"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		log.Fatalf("need one arg")
	}

	wr, err := parser.NewWasmReader(args[0])
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	defer wr.Close()

	wr.ReadHeader()
	wr.ReadVer()

	parser.Section(wr)
	parser.Section(wr)
	parser.Section(wr)
	parser.Section(wr)
	parser.Section(wr)
	parser.Section(wr)
	return
}
