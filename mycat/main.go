package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func loadData(path string) ([]string, error) {
	f, err := os.Open(path)
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

	// line number
	var ln int

	for _, path := range flag.Args() {
		rows, err := loadData(path)
		if err != nil {
			fmt.Println(err)
			break
		}
		for _, r := range rows {
			if *number {
				ln++
				fmt.Printf("%d: %s\n", ln, r)
			} else {
				fmt.Println(r)
			}
		}
	}
}
