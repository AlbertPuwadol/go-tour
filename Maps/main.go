package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	stringList := strings.Fields(s)
	stringMap := make(map[string]int)
	for _, v := range stringList {
		stringMap[v]++
	}
	return stringMap
}

func main() {
	wc.Test(WordCount)
}
