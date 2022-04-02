package main

import (
	"fmt"
	"strings"
)

func camelCaseToWords(str string) []string {
	var ret []string

	// total := 0
	count := 0
	prev_count := count
	for count < len(str) - 1 {
		for i := count + 1; i < len(str); i++ {
			if str[i] >= 65 && str[i] <= 90 || i == len(str) - 1 {
				prev_count = count
				count = i
				// total += count
				if count == len(str) - 1 {
					ret = append(ret, str[prev_count:])
				} else {
					ret = append(ret, str[prev_count:count])
				}
				break
			}
		}
	}

	return ret
}

func camelCaseToWords2(str string) []string {
	var ret []string

	for i, ch := range str {
		fmt.Printf("%#U starts byte at position: %d\n", ch, i)
	}

	return ret
}

func camelCaseToWords3(str string) []string {
	var ret []string

	prev := 0
	for i, ch := range str {
		s := string(ch)
		if strings.ToUpper(s) == s {
			ret = append(ret, str[prev:i])
			prev = i
		}
	}
	ret = append(ret, str[prev:])

	return ret
}

func main() {
	fmt.Println(camelCaseToWords3("saveChangesInTheEditor"))
	fmt.Println(camelCaseToWords2("５時半に学校で音楽を聞きます。"))
}
