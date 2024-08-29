package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

func main() {
	help := flag.Bool("help", false, "Show help message")
	wordCount := flag.Int("w", 100, "Number of maximum words")
	startPrefix := flag.String("p", "", "Starting Prefix")
	prefixLen := flag.Int("l", 2, "Prefix length")
	flag.Parse()

	if *help {
		PrintHelp()
		return
	}

	if *wordCount <= 0 || *wordCount > 10000 {
		fmt.Fprintln(os.Stderr, "Error: Invalid word count")
		os.Exit(1)

	}

	if *prefixLen <= 0 || *prefixLen > 5 {
		fmt.Fprintln(os.Stderr, "Incorrect prefix length")
		os.Exit(1)
	}

	words := HandleStdin()

	if len(words) == 0 {
		fmt.Fprintln(os.Stderr, "Error: Text is empty!")
		os.Exit(1)
	}

	MarkovDictionary := make(map[string][]string)

	var firstPrefix string

	for i := 0; i < len(words)-*prefixLen; i++ {
		prefix := strings.Join(words[i:i+*prefixLen], " ")
		suffix := words[i+*prefixLen]
		MarkovDictionary[prefix] = append(MarkovDictionary[prefix], suffix)
		if i == 0 {
			firstPrefix = prefix
		}
	}

	if len(MarkovDictionary) == 0 {
		fmt.Fprintln(os.Stderr, "Error: markov dictionary is empty")
		os.Exit(1)
	}

	var generatedText string

	if *startPrefix != "" {
		if ValidStartingPrefix(*startPrefix, MarkovDictionary, *prefixLen) {
			generatedText = MarkovAlgorithm(MarkovDictionary, *prefixLen, *wordCount, *startPrefix)
		} else {
			fmt.Fprintln(os.Stderr, "Error: needed prefix is not found in text")
			os.Exit(1)
		}
	} else {
		if ValidStartingPrefix(firstPrefix, MarkovDictionary, *prefixLen) {
			generatedText = MarkovAlgorithm(MarkovDictionary, *prefixLen, *wordCount, firstPrefix)
		} else {
			fmt.Fprintln(os.Stderr, "Error: the default starting prefix is invalid")
			os.Exit(1)
		}
	}

	fmt.Println(generatedText)
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

func ValidStartingPrefix(startingPrefix string, MarkovDictionary map[string][]string, prefixLen int) bool {
	_, exists := MarkovDictionary[startingPrefix]
	return exists
}

func MarkovAlgorithm(MarkovDictionary map[string][]string, prefixLen int, length int, startPrefix string) string {
	var sb strings.Builder
	var prefix string

	if startPrefix != "" {
		prefix = startPrefix
	} else {
		prefixes := make([]string, 0, len(MarkovDictionary))
		for p := range MarkovDictionary {
			prefixes = append(prefixes, p)
		}
		prefix = prefixes[rand.Intn(len(prefixes))]
	}

	sb.WriteString(prefix)
	words := strings.Split(prefix, " ")
	wordCount := len(words)

	for wordCount < length {
		suffixes, exists := MarkovDictionary[prefix]
		if !exists || len(suffixes) == 0 {
			break
		}

		suffix := suffixes[rand.Intn(len(suffixes))]
		sb.WriteString(" ")
		sb.WriteString(suffix)

		words = append(words[1:], suffix)
		prefix = strings.Join(words, " ")

		wordCount++
	}

	return sb.String()
}

func PrintHelp() {
	fmt.Println(`Markov Chain text generator.

	Usage:
	  markovchain [-w <N>] [-p <S>] [-l <N>]
	  markovchain --help
	
	Options:
	  --help  Show this screen.
	  -w N    Number of maximum words
	  -p S    Starting prefix
	  -l N    Prefix length`)
}
