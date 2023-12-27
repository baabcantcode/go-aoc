package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// / also times cuz why not
func Scan(fname string, f func(*bufio.Scanner) string) string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	start := time.Now()
	d := f(scanner)
	end := time.Since(start)
	fmt.Println("process took", end)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return d
}
