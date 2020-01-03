package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type WasmReader struct {
	rdr *bufio.Reader
	file *os.File
}

func NewWasmReader(fn string) (*WasmReader, error) {
	file, err := os.Open(fn)
	if err != nil {
		log.Fatalf("Error occured %s", err)
		return nil, err
	}

	rdr := bufio.NewReader(file)
	wr := WasmReader{file: file, rdr: rdr}

	return &wr, nil
}

func (wr *WasmReader) ReadHeader() []byte {
	hdr := make([]byte, 4)
	n, err := wr.rdr.Read(hdr)
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

	return hdr
}

func (wr *WasmReader) ReadVer() []byte {
	ver := make([]byte, 4)
	n, err := wr.rdr.Read(ver)
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	if n < 4 {
		log.Fatalf("File too short!")
	}

	return ver
}

func (wr *WasmReader) ReadType() int {
	b, err := wr.rdr.ReadByte()
	if err != nil {
		log.Fatalf("Error occured %s", err)
	}
	return int(b)
}

func (wr *WasmReader) ReadU32() uint32 {
	return U32(wr.rdr)
}

func (wr *WasmReader) ReadName() string {
	return Name(wr.rdr)
}

func (wr *WasmReader) ReadLimit() {
	Limit(wr.rdr)
}

func (wr *WasmReader) Reader() *bufio.Reader {
	return wr.rdr
}

func (wr *WasmReader) Close() {
	wr.file.Close()
}
