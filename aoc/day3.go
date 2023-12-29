package aoc

import (
	util "aoc2023/util"
	"bufio"
	"fmt"

	"github.com/spf13/cobra"
)

var day3cmd = &cobra.Command{
	Use: "3",
	Run: func(cmd *cobra.Command, args []string) {
		util.Scan("./data/d3p1.txt", d3part1)
		// util.Scan("./data/d2p1.txt", d2part2)
	},
}

func init() {
	aocCmd.AddCommand(day3cmd)
}

func d3part1(scanner *bufio.Scanner) string {
	sum := 0
	prevstr, curstr, nextstr := "", "", ""
	for scanner.Scan() {
		// just gonna trust its ascii, since i know the input
		nextstr = scanner.Text()
		print("PREV", prevstr, "CUR", curstr, "NEXT", nextstr)
		// do logic here
		prevstr = curstr
		curstr = nextstr
	}
	// do final check on nextstr adjacencies here
	print(sum)
	return fmt.Sprint(sum)
}
