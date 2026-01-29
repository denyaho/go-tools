package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	f, err := os.Open("test.txt")
	defer f.Close()
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	n := 10
	buffer := make([]string, n)
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		buffer[count%n] = scanner.Text()
		count++
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	displayCount := min(n, count)
	startIndex := count%n
	if count < n {
		startIndex = 0
	}
	for i :=0 ;i < displayCount; i++ {
		idx := (startIndex +i)% n
		fmt.Println(buffer[idx])
	}
}
