package aoc

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var print = fmt.Println

var aocCmd = &cobra.Command{
	Use: "aoc",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := aocCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)

		os.Exit(1)

	}
}
