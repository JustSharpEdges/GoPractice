package main

import "fmt"

func findLongestString(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	longest := strings[0]
	for _, s := range strings[1:] {
		if len(s) > len(longest) {
			longest = s
		}
	}

	return longest
}

func main() {
	fmt.Println(findLongestString("cat", "dog", "elephant"))
	fmt.Println(findLongestString("apple", "banana", "pear"))
	fmt.Println(findLongestString("one"))
	fmt.Println(findLongestString())
}
