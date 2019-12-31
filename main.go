package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"net.fujlog/go-wasm/parser"
)

func main() {
	file, err := os.Open("./wasm/a.out.wasm")
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	defer file.Close()
	/*stats, statsErr := file.Stat()
	 if statsErr != nil {
		log.Fatalf("Error occured %s", statsErr)
	 }
	 var size int64 = stats.Size()
	bytes := make([]byte, size)
	*/
	bufr := bufio.NewReader(file)
	var n int
	hdr := make([]byte, 4)
	ver := make([]byte, 4)
	n, err = bufr.Read(hdr)
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	if n < 4 {
		log.Fatalf("File too short!")
	}
	if !(hdr[0] == 0 && string(hdr[1:]) == "asm") {
		log.Fatalf("Not a wasm file.")
	}
	fmt.Println(string(hdr[1:]))
	n, err = bufr.Read(ver)
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	if n < 4 {
		log.Fatalf("File too short!")
	}
	parser.Section(bufr)
	parser.Section(bufr)
	parser.Section(bufr)
	parser.Section(bufr)
	parser.Section(bufr)
	parser.Section(bufr)
	return
}
