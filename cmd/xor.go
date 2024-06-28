package cmd

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var isHex bool
var outFormat string

func init() {
	rootCmd.PersistentFlags().BoolVar(&isHex, "hex", false, "specify if entry value is hex")
	rootCmd.PersistentFlags().StringVar(&outFormat, "out", "hex", "set the output format. hex | ascii")
}

var rootCmd = &cobra.Command{
	Use: "",
	Run: func(cmd *cobra.Command, args []string) {
		var buf1 []byte
		var buf2 []byte
		var err error

		if isHex {
			buf1, err = hex.DecodeString(args[0])
			if err != nil {
				log.Fatal("failed to decode hex1: " + err.Error())
			}

			buf2, err = hex.DecodeString(args[1])
			if err != nil {
				log.Fatal("failed to decode hex2: " + err.Error())
			}
		} else {
			buf1 = []byte(args[0])
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
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
