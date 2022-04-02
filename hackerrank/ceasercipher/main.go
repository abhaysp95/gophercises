package main

import (
	"fmt"
	"strings"
)

func ceaserchiper(str string, k int) string {
	ret := make([]byte, 0, len(str))

	for _, ch := range str {
		s := string(ch)
		if (ch < 65 || ch > 90) && (ch < 97 || ch > 122) {
			ret = append(ret, byte(int(ch)))
		} else if strings.ToUpper(s) == s {
			ret = append(ret, byte((int(ch) - 'A' + k) % 26) + 'A')
		} else {
			fmt.Printf("%c: %d, %d\n", ch, ch, (int(ch) + k) % 122)
			ret = append(ret, byte((int(ch) - 'a' + k) % 26) + 'a')
		}
	}

	return string(ret)
}

const (
	alphabetLower = "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"  // or use strings.ToUpper
)

// somewhat generic approach
func ceaserchiper2(input string, k int) string {
	var ret string

	for _, ch := range input {
		switch {
		case strings.IndexRune(alphabetUpper, ch) >= 0:
			ret += string(rotate(ch, k, []rune(alphabetUpper)))
		case strings.IndexRune(alphabetLower, ch) >= 0:
			ret += string(rotate(ch, k, []rune(alphabetLower)))
		default:
			ret += string(ch)
		}
	}

	return ret
}

func rotate(r rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), r)
	if idx < 0 {
		panic("idx < 0")
	}
	return key[(idx + delta) % len(key)]
}

func main() {
	fmt.Println(ceaserchiper("middle-Outz", 2))
}
