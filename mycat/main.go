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
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rows, nil
}

func main() {
	var number = flag.Bool("n", false, "display line numbers")
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
				fmt.Fprintf(os.Stdout, "%d: %s\n", lineNumber, r)
			} else {
				fmt.Fprintln(os.Stdout, r)
			}
		}
	}
}
