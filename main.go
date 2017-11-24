package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
)

// Pair k/v for sorting
type Pair struct {
	Key   string
	Value int
}

// PairList list of Pair
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(p))
	return p
}

func main() {
	var whitespace = regexp.MustCompile(`\s`)
	var data []byte
	var err error
	var tokenCount = make(map[string]int)
	var topN = os.Args[1]
	data, err = ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Data: ", string(data))
	document := whitespace.Split(string(data), -1)

	fmt.Println("Tokens: ", document)

	for _, token := range document {
		if token == "" {
			continue
		}

		// if the token exists, increment its count
		if _, ok := tokenCount[token]; ok {
			tokenCount[token]++
		} else {
			tokenCount[token] = 1
		}
	}

	fmt.Println("Tokens and the number of their occurrences ", tokenCount)

	// display the top N occurences
	if s, err := strconv.ParseInt(topN, 10, 32); err == nil {
		sorted := sortMapByValue(tokenCount)
		i := 0
		for _, pair := range sorted {
			if int64(i) >= s {
				return
			}
			fmt.Println(pair.Key, ": ", pair.Value)
			i++
		}

	}

}
