package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	var whitespace = regexp.MustCompile(`\s`)
	var data []byte
	var err error
	var tokenCount = make(map[string]int)
	data, err = ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Data ", string(data))
	document := whitespace.Split(string(data), -1)

	fmt.Println("tokens: ", document)

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
}
