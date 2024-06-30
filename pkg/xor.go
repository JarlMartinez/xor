package xor

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Perofm(a, b string) (*[]uint8, error) {
	buf1, err := parseInput(a)
	if err != nil {
		return nil, fmt.Errorf("parsing input a: %e", err)
	}

	buf2, err := parseInput(b)
	if err != nil {
		return nil, fmt.Errorf("parsing input b: %w", err)
	}

	out := make([]uint8, len(buf1))
	for i := range buf1 {
		out[i] = buf1[i] ^ buf2[i%len(buf2)]
	}

	return &out, nil
}

func parseInput(in string) ([]byte, error) {
	if isHex(&in) {
		buf, err := hex.DecodeString(in[2:]) // trim out 0x
		if err != nil {
			return nil, fmt.Errorf("decoding as hex: %w", err)
		}
		return buf, nil
	}
	// see if it's base64
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(in)))
	_, err := base64.StdEncoding.Decode(buf, []byte(in))
	if err == nil {
		return buf, nil
	}
	if _, ok := err.(base64.CorruptInputError); ok {
		// treat as plaintext
		return []byte(in), nil
	} else {
		return nil, fmt.Errorf("decoding as base64: %w", err)
	}
}

func isHex(in *string) bool {
	return (*in)[0] == '0' && (*in)[1] == 'x'
}
