package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type LogResult struct {
	Details  map[string]int
	Category map[int]int
}

func analyze_log(r io.Reader) (LogResult, error) {
	detailed := make(map[string]int)
	category := make(map[int]int)
	const statusIndex = 8
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) <= statusIndex {
			continue
		}
		statusStr := fields[statusIndex]
		detailed[statusStr]++

		if code, err := strconv.Atoi(statusStr); err == nil {
			category[code/100*100]++
		}
	}
	return LogResult{Details: detailed, Category: category}, nil
}

func main() {
	result, err := analyze_log(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error analyzing log: %v\n", err)
		os.Exit(1)
	}
	keys := make([]string, 0, len(result.Details))
	for k := range result.Details {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, result.Details[k])
	}

	categorykeys := make([]int, 0, len(result.Category))
	for k := range result.Category {
		categorykeys = append(categorykeys, k)
	}
	sort.Ints(categorykeys)
	for _, category := range categorykeys {
		fmt.Printf("%dxx: %d\n", category/100, result.Category[category])
	}

}
