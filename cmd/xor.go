package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var xorCmd = &cobra.Command{
	Use: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello cobra")
	},
}
