package aoc

import (
	util "aoc2023/util"
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

var day1cmd = &cobra.Command{
	Use: "1",
	Run: func(cmd *cobra.Command, args []string) {
		util.Scan("./data/d1p1.txt", d1)
		util.Scan("./data/d1p1.txt", d1p2)
	},
}

func init() {
	aocCmd.AddCommand(day1cmd)
}

var regex = regexp.MustCompile(`^(one|two|three|four|five|six|seven|eight|nine|[0-9])`)

var translate_map = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4",
	"five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

func d1p2(scanner *bufio.Scanner) string {
	sum := 0
	for scanner.Scan() {
		fullstr := scanner.Text()
		runedstr := []rune(fullstr)
		lenstr := len(runedstr)
		var lstr, rstr string
		for i := range runedstr {
			match := regex.FindString(string(runedstr[i:min(i+5, lenstr)]))
			if match != "" {
				lstr = translate_map[match]
				if lstr == "" {
					lstr = match
				}
				// fmt.Println(lstr, match, fullstr)
				break
			}
		}
		for i := lenstr - 1; i >= 0; i-- {
			match := regex.FindString(string(runedstr[i:min(i+5, lenstr)]))
			if match != "" {
				rstr = translate_map[match]
				if rstr == "" {
					rstr = match
				}
				// fmt.Println(rstr, match, fullstr)
				break
			}

		}
		if lstr == "" {
			continue
		}
		i, err := strconv.Atoi(string(lstr) + string(rstr))
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(fullstr, string(lstr), string(rstr), i)
		sum += i

	}
	fmt.Println(sum)
	return fmt.Sprint(sum)
}

func d1(scanner *bufio.Scanner) string {
	sum := 0
	for scanner.Scan() {
		fullstr := scanner.Text()
		runedstr := []rune(fullstr)
		lstr, rstr := '-', '-'
		for _, v := range runedstr {
			if v > 47 && v < 58 {
				lstr = v
				break
			}
		}
		if lstr == '-' {
			continue
		}
		for i := len(runedstr) - 1; i >= 0; i-- {
			v := runedstr[i]
			if v > 47 && v < 58 {
				rstr = v
				break
			}
		}

		i, err := strconv.Atoi(string(lstr) + string(rstr))
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(fullstr, string(lstr), string(rstr), i)
		sum += i
	}
	fmt.Println(sum)
	return fmt.Sprint(sum)
}
