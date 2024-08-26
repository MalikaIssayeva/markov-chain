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

	prefixLen := 2
	if prefixLen < 1 || prefixLen > 5 {
		fmt.Fprintln(os.Stderr, "Error: Prefix length must be between 1 and 5.")
		os.Exit(1)
	}
	MarkovDictionary := make(map[string][]string)
	for i := 0; i < len(words)-prefixLen; i++ {
		prefix := strings.Join(words[i:i+prefixLen], " ")
		suffix := words[i+prefixLen]
		MarkovDictionary[prefix] = append(MarkovDictionary[prefix], suffix)
	}
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

func MarkovAlgoritm(MarkovDictionary map[string][]string, prefixLen int, length int) string {
	prefixes := make([]string, 0, len(MarkovDictionary))
	for prefix := range MarkovDictionary {
		prefixes = append(prefixes, prefix)
	}
	if len(prefixes) == 0 {
		// fmt.Fprintln()
	}
	return "lll"
}

// fmt.Println(MarkovDictionary)
// wordCount := flag.Int("w", 100, "Number of maximum words")
// flag.Parse()
// if *wordCount <= 0 || *wordCount > 10000 {
// 	fmt.Fprintln(os.Stderr, "Error: Invalid word count")
// 	os.Exit(1)
// }
