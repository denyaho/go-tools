package main

import (
	"fmt"
	"os"
	"bytes"
	"bufio"
	"flag"
	"strings"
	"encoding/json"
)

type Data struct {
	Line int `json:"line"`
	Text string `json:"text"`
}

func main(){
	isJson := flag.Bool("json", false, "Output in JSON format")
	isNumbering := flag.Bool("n", false, "Number the output lines")
	separator := flag.String("s", "\n", "Separator to use between lines")
	skipEmpty := flag.Bool("skip-empty", false, "Skip empty lines")
	flag.Parse()

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i:= bytes.Index(data, []byte(*separator)); i>= 0{
			return i+ len(*separator), data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	
	for scanner.Scan() {
		text := scanner.Text()
		var line string

		if *separator != "\n" {
			cleanText := strings.ReplaceAll(text, "\n", "")
			cleanText = strings.ReplaceAll(cleanText, "\r", "")
			line = strings.TrimSpace(cleanText)
		}else{
			line = strings.TrimSpace(text)
		}
		if *skipEmpty && len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading standard input:", err)
		os.Exit(1)
	}

	for i, line := range lines {
		if *isJson{
			data := Data{Line: i + 1, Text: line}
			v, err := json.Marshal(data)
			if err != nil{
				fmt.Fprintln(os.Stderr, "JSON Error")
				continue
			}
			fmt.Println(string(v))
		}else if *isNumbering {
			fmt.Printf("%d: %s\n", i+1, line)
		}else {
			fmt.Println(line)
		}
	}
}