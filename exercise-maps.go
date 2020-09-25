package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	data := make(map[string]int)

	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {

		v, ok := data[words[i]]
		if ok {
			data[words[i]] = v + 1
		} else {
			data[words[i]] = 1
		}

	}
	return data

}

func main() {
	wc.Test(WordCount)
}
