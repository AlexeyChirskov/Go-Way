package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

const minStrSize = 3

func main() {
	fmt.Println(comma("123"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("1234524.35234"))

	fmt.Println(isAnagram("aaaa", "aaaa"))
	fmt.Println(isAnagram("aaaa", "aaaazz"))
	fmt.Println(isAnagram("abba", "baab"))
	fmt.Println(isAnagram("ab🐻ba", "baab🐻"))
}

func comma(s string) string {
	n := strings.Index(s, ".")
	if n == -1 {
		n = len(s)
	}

	if n <= 3 {
		return s
	}

	var buf bytes.Buffer

	i := 0
	for j := 0; j < n; j++ {
		if j%3 == 0 && j > 1 {
			if buf.Len() > 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(s[j-3 : j])
			i = j
		}
	}
	buf.WriteString(", " + s[i:])

	return buf.String()
}

func isAnagram(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}

	if len(s1) != len(s2) {
		return false
	}

	m1 := make(map[rune]int64)
	m2 := make(map[rune]int64)

	for _, v := range s1 {
		m1[v]++
	}

	for _, v := range s2 {
		m2[v]++
	}

	return reflect.DeepEqual(m1, m2)
}
