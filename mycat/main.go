package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func readFile(fn string) ([]string, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	return rows, nil
}

func main() {
	var number = flag.Bool("n", false, "add line number")
	flag.Parse()
	fileNames := flag.Args()

	var lineNumber int
	for _, fn := range fileNames {
		rows, err := readFile(fn)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, r := range rows {
			if *number {
				lineNumber++
				fmt.Printf("%d: %s\n", lineNumber, r)
			} else {
				fmt.Println(r)
			}
		}
	}
}
