package main

/**
709. To Lower Case

Given a string s, return the string after replacing every uppercase letter with the same lowercase letter.

Example 1:

Input: s = "Hello"
Output: "hello"

Example 2:

Input: s = "here"
Output: "here"

Example 3:

Input: s = "LOVELY"
Output: "lovely"


*/
import (
	"fmt"
	"strings"
)

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func main() {
	fmt.Println(toLowerCase("Hello"))
}
