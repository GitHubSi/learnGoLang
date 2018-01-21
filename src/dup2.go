package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	count := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLine(os.Stdin, count)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}

			countLine(f, count)
			f.Close()
		}
	}

	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLine(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]  ++
	}
}
