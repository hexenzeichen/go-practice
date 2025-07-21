package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)

	sliceOfStrings := strings.Fields(s)

	for _, value := range sliceOfStrings {

		_, ok := m[value]
		if ok {
			m[value] = m[value] + 1
		} else {
			m[value] = 1
		}

	}

	return m
}

func main() {
	wc.Test(WordCount)
}
