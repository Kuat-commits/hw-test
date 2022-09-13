package main

import "fmt"

// this is a comment

func main() {
	s := "Hello, OTUS!"
	s = Reverse(s)
	fmt.Println(s)
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
