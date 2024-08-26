package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("the_great_gatsby.txt")
	if err != nil {
		fmt.Printf("There is a problem with reading a file: %v\n", err)
		return
	}
	defer file.Close()
	read := bufio.NewReader(file)
	text, err := io.ReadAll(read)
	if err != nil {
		fmt.Printf("There is a problem with reading a text: %v\n", err)
		os.Exit(1)
	}
	words := strings.Fields(string(text))
	if len(words) == 0 {
		fmt.Println("Text is empty!")
		os.Exit(1)
	}
	MarkovDictionary := make(map[string][]string)
	prefixLen := 2
	for i := 0; i < len(words)-prefixLen; i++ {
		prefix := strings.Join(words[i:i+prefixLen], " ")
		suffix := words[i+prefixLen]
		MarkovDictionary[prefix] = append(MarkovDictionary[prefix], suffix)
	}
}
