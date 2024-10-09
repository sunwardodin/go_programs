package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	// Open a file
	f, err := os.Open("gatsby.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()

	words, err := freq(f)
	if err != nil {
		log.Fatalf("error from freq in main: %s", err)
	}

	pairs := sortWordFrequency(words)

	for _, pair := range pairs {
		fmt.Printf("%s \t\t %d\n", pair.Key, pair.Value)
	}

	w, n, err := maxWord(words)
	if err != nil {
		log.Fatalf("error with maxWord: %s\n", err)
	}

	fmt.Printf("%#v has the highest frequency of %d\n", w, n)
}

func freq(r io.Reader) (map[string]int, error) {
	wordFreq := make(map[string]int)

	s := bufio.NewScanner(r)
	s.Split((bufio.ScanWords))

	for s.Scan() {
		word := strings.ToLower(s.Text())
		wordFreq[word]++
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return wordFreq, nil
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortWordFrequency(m map[string]int) PairList {
	pairs := make(PairList, len(m))
	i := 0
	for key, value := range m {
		pairs[i] = Pair{key, value}
		i++
	}

	sort.Sort(pairs)

	return pairs
}

func maxWord(m map[string]int) (string, int, error) {
	if len(m) == 0 {
		return "", 0, fmt.Errorf("empty map")
	}

	maxW := ""
	maxF := 0

	for k, v := range m {
		if v > maxF {
			maxW = k
			maxF = v
		}
	}
	return maxW, maxF, nil
}
