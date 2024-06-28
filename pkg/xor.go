package xor

import (
	"encoding/hex"
	"fmt"
)

func Perofm(a, b string) (*[]uint8, error) {
	var buf1 []byte
	var buf2 []byte
	var err error

	if isHex(&a) {
		buf1, err = hex.DecodeString(a[2:]) // trim out 0x
		if err != nil {
			return nil, fmt.Errorf("decoding hex a input: %s", err)
		}
	} else {
		// treat as plaintext
		buf1 = []byte(a)
	}

	if isHex(&b) {
		buf2, err = hex.DecodeString(b[2:]) // trim out 0x
		if err != nil {
			return nil, fmt.Errorf("decoding hex b input: %s", err)
		}
	} else {
		// treat as plaintext
		buf2 = []byte(b)
	}

	out := make([]uint8, len(buf1))
	for i := range buf1 {
		out[i] = buf1[i] ^ buf2[i%len(buf2)]
	}

	return &out, nil
}

func isHex(in *string) bool {
	return (*in)[0] == '0' && (*in)[1] == 'x'
}
