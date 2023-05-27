package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/quic-go/qpack"
)

func Chunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}

func HexConvertDec(chunk []string) []byte {
	var decChunk []byte
	for i := range chunk {
		bs, _ := hex.DecodeString(chunk[i])
		decChunk = append(decChunk, bs[0])
	}
	return decChunk
}

func main() {
	var parseString string
	flag.StringVar(&parseString, "s", "", "the hexadecimal content that need to be translated")
	flag.Parse()

	if len(parseString) == 0 {
		fmt.Println("please input what information you want to translated")
	}

	splitedString := Chunks(parseString, 2)
	decChunk := HexConvertDec(splitedString)

	decoder := qpack.NewDecoder(nil)
	hf, err := decoder.DecodeFull(decChunk)
	if err != nil {
		fmt.Println(err)
	}
	for i := range hf {
		fmt.Printf("%s: %s\r\n", hf[i].Name, hf[i].Value)
	}
}
