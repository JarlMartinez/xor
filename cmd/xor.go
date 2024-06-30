package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	xor "github.com/JarlMartinez/xor/pkg"
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

		stdi, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			log.Fatal("failed reading stdin: ", err.Error())
		}

		var in string

		if len(stdi) > 0 {
			in = string(stdi)
		} else {
			in = args[0]
		}

		buf, err := xor.Perofm(in, args[1])
		if err != nil {
			log.Fatal("failed to perform xor: " + err.Error())
		}

		switch outFormat {
		case "hex":
			fmt.Printf("%x\n", buf)
		case "ascii":
			fmt.Printf("%s\n", string(*buf))
		case "all":
			fmt.Printf("\thex:   %0x\n", *buf)
			fmt.Printf("\tascii:  %s\n", string(*buf))
		}
	},
}

func isHex(in *string) bool {
	return (*in)[0] == '0' && (*in)[1] == 'x'
}
