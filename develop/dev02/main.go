package main

import (
	"fmt"
	"strings"
	"unicode"
)

func unpack(s string) string {
	ans := ""
	seq := []rune(s)
	l := 0
	r := 1
	for r < len(seq) {
		if unicode.IsDigit(seq[r]) {
			if unicode.IsDigit(seq[l]) {
				return ""
			} else {
				ans += strings.Repeat(string(seq[l]), int(seq[r])-'0')
			}
		} else if !unicode.IsDigit(seq[l]) {
			ans += string(seq[l])
		}
		l++
		r++
	}
	fmt.Println(l, r)
	if unicode.IsDigit(seq[l]) {
		return ans
	}
	return ans + string(seq[l])
}

func main() {
	var inp string
	fmt.Print("Input string: ")
	fmt.Scan(&inp)
	fmt.Println(unpack(inp))
}
