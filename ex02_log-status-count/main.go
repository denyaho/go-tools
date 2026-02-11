package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){
	var status_count []string
	var status_map map[int]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		split_text := strings.Split(text, " ")
		if strings.HasPrefix(text, "http") || strings.HasPrefix(text, "https") {
			
		}		
	}


}