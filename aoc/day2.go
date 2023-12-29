package aoc

import (
	util "aoc2023/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"

	// "fmt"
	// "log"
	// "regexp"
	// "strconv"

	"github.com/spf13/cobra"
)

var day2cmd = &cobra.Command{
	Use: "2",
	Run: func(cmd *cobra.Command, args []string) {
		util.Scan("./data/d2p1.txt", d2part1)
		util.Scan("./data/d2p1.txt", d2part2)
	},
}

func init() {
	aocCmd.AddCommand(day2cmd)
}

var rcubes = 12
var bcubes = 14
var gcubes = 13

func d2part2(scanner *bufio.Scanner) string {
	sum := 0
	for scanner.Scan() {
		// just gonna trust its ascii, since i know the input
		fullstr := scanner.Text()
		tag_split := strings.Split(fullstr, ":")
		if len(tag_split) < 2 {
			continue
		}
		maxr, maxb, maxg := 0, 0, 0
		pulls := strings.Split(tag_split[1], ";")
		for _, ball_counts := range pulls {
			iballs := strings.Split(ball_counts, ",")
			// print(iballs)
			for _, ball_count := range iballs {
				count_color := strings.Split(strings.Trim(ball_count, " "), " ")
				if len(count_color) < 2 {
					continue
				}
				count, err := strconv.Atoi(count_color[0])
				if err != nil {
					print("bad conversion", count_color[0])
					continue
				}
				switch color := strings.Trim(count_color[1], " "); color {
				case "red":
					maxr = max(count, maxr)
				case "blue":
					maxb = max(count, maxb)
				case "green":
					maxg = max(count, maxg)
				}

			}
		}
		sum += maxg * maxr * maxb
		// print(fullstr, maxr, maxg, maxb, sum)
	}
	print(sum)
	return fmt.Sprint(sum)
}

func d2part1(scanner *bufio.Scanner) string {
	sum := 0
	for scanner.Scan() {
		// just gonna trust its ascii, since i know the input
		fullstr := scanner.Text()
		tag_split := strings.Split(fullstr, ":")
		if len(tag_split) < 2 {
			continue
		}
		invalid := false
		pulls := strings.Split(tag_split[1], ";")
		for _, ball_counts := range pulls {
			iballs := strings.Split(ball_counts, ",")
			// print(iballs)
			for _, ball_count := range iballs {
				count_color := strings.Split(strings.Trim(ball_count, " "), " ")
				if len(count_color) < 2 {
					continue
				}
				count, err := strconv.Atoi(count_color[0])
				if err != nil {
					print("bad conversion", count_color[0])
					continue
				}
				switch color := strings.Trim(count_color[1], " "); color {
				case "red":
					if count > rcubes {
						invalid = true
						break
					}
				case "blue":
					if count > bcubes {
						invalid = true
						break
					}
				case "green":
					if count > gcubes {
						invalid = true
						break
					}
				}

			}
		}
		if invalid == false {
			// print(fullstr, "valid")
			count, err := strconv.Atoi(strings.TrimPrefix(tag_split[0], "Game "))
			if err != nil {
				continue
			}
			sum += count
		} else {
			// print(fullstr, "invalid")
		}
	}
	print(sum)
	return fmt.Sprint(sum)
}
