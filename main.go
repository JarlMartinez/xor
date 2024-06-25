package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {

	hex1 := os.Args[1]
	hex2 := os.Args[2]

	buf1, err := hex.DecodeString(hex1)
	if err != nil {
		log.Fatal("failed to decode hex1: " + err.Error())
	}

	buf2, err := hex.DecodeString(hex2)
	if err != nil {
		log.Fatal("failed to decode hex2: " + err.Error())
	}

	out := make([]uint8, len(buf1))
	for i := range buf1 {
		out[i] = buf1[i] ^ buf2[i]
	}

	fmt.Printf("%x\n", out)
}
