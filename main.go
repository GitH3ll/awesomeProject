package main

import "fmt"

func main() {
	m := countRunes("aaaaaaasgdfgdgbbbb")
	for k, v := range m {
		fmt.Println(string(k), " ", v)
	}
}

// "aaaaaabbb"
func countRunes(s string) map[rune]int {
	m := make(map[rune]int)

	for _, r := range s {
		_, ok := m[r]
		if ok {
			m[r] = m[r] + 1
			continue
		}
		m[r] = 1
	}
	return m
}
