package cmd

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var outFormat string

func init() {
	rootCmd.PersistentFlags().StringVar(&outFormat, "out", "all", "set the output format. hex | ascii | all")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use: "",
	Run: func(cmd *cobra.Command, args []string) {
		var buf1 []byte
		var buf2 []byte
		var err error

		if isHex(&args[0]) {
			buf1, err = hex.DecodeString(args[0][2:]) // trim out 0x
			if err != nil {
				log.Fatal("failed to decode arg1 as hex: " + err.Error())
			}
		} else {
			// treat as plaintext
			buf1 = []byte(args[0])
		}

		if isHex(&args[1]) {
			buf2, err = hex.DecodeString(args[1][2:]) // trim out 0x
			if err != nil {
				log.Fatal("failed to arg2 as hex: " + err.Error())
			}
		} else {
			// treat as plaintext
			buf2 = []byte(args[1])
		}

		out := make([]uint8, len(buf1))
		for i := range buf1 {
			out[i] = buf1[i] ^ buf2[i%len(buf2)]
		}

		switch outFormat {
		case "hex":
			fmt.Printf("%x\n", out)
		case "ascii":
			fmt.Printf("%s\n", string(out))
		case "all":
			fmt.Printf("\thex:   %0x\n", out)
			fmt.Printf("\tascii:  %s\n", string(out))
		}
	},
}

func isHex(in *string) bool {
	return (*in)[0] == '0' && (*in)[1] == 'x'
}
