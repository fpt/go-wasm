package main

import (
	"log"

	"net.fujlog/go-wasm/parser"
)

func main() {
	wr, err := parser.NewWasmReader("./wasm/a.out.wasm")
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
