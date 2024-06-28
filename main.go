package main

import (
	"encoding/hex"
	"fmt"

	"github.com/JarlMartinez/xor/cmd"
)

func main() {
	cmd.Execute()
	// test()
}

func test() {
	in := "ICE"
	fmt.Println([]byte(in))
	fmt.Println(hex.EncodeToString([]byte(in)))
	fmt.Println([]byte(hex.EncodeToString([]byte(in))))
}

// hex:     4275726e696e672027656d2c20696620796f752061696e277420717569636b20616e64206e696d626c650a4920676f206372617a79207768656e2049206865617220612063796d62616c
// bytes:   [66 117 114 110 105 110 103 32 39 101 109 44 32 105 102 32 121 111 117 32 97 105 110 39 116 32 113 117 105 99 107 32 97 110 100 32 110 105 109 98 108 101 10 73 32 103 111 32 99 114 97 122 121 32 119 104 101 110 32 73 32 104 101 97 114 32 97 32 99 121 109 98 97 108]
// ascii:   Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal
