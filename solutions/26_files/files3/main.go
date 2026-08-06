// Concept: bufio.Scanner reads input line by line
// Task: open numbers.txt, scan it line by line, and print the sum of the numbers
// Expected output: sum: 6
// Hint: sc := bufio.NewScanner(f); for sc.Scan() { ... strconv.Atoi(sc.Text()) ... };
//       check sc.Err() after the loop, and defer f.Close() right after opening.
//       Run from this directory (`cd exercises/26_files/files3 && go run .`) or
//       verify with `go test ./exercises/26_files/files3` (Go doc: bufio)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("numbers.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	sum := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		n, err := strconv.Atoi(sc.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += n
	}
	if err := sc.Err(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("sum:", sum)
}
