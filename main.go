package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	words := HandleStdin()
	if len(words) == 0 {
		fmt.Fprintln(os.Stderr, "Error: Text is empty!")
		os.Exit(1)
	}
	MarkovDictionary := make(map[string][]string)
	prefixLen := 2
	for i := 0; i < len(words)-prefixLen; i++ {
		prefix := strings.Join(words[i:i+prefixLen], " ")
		suffix := words[i+prefixLen]
		MarkovDictionary[prefix] = append(MarkovDictionary[prefix], suffix)
	}
	fmt.Println(MarkovDictionary)
}

func HandleStdin() []string {
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: incorrect input")
		os.Exit(1)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Fprintln(os.Stderr, "Error: no input text")
		os.Exit(1)
	}

	buf := new(strings.Builder)
	io.Copy(buf, os.Stdin)
	input := buf.String()
	return strings.Fields(input)
}
