package main

import (
	"fmt"
)

func main() {
	ptr := "aaba"
	str := "aabaabaa"
	k := len(ptr)
	i := 0
	j := 0
	maps := map[string]int{
		"a": 3,
		"b": 1,
	}

	count := 2
	ans := 0

	for j < len(str) {

		maps[string(str[j])]--
		if maps[string(str[j])] == 0 {
			count--
		}
		if (j - i + 1) < k {
			j++
		} else if (j - i + 1) == k {
			if count == 0 {
				ans++
			}
			maps[string(str[i])]++
			if maps[string(str[i])] == 1 {
				count++
			}
			i++
			j++
		}
	}
	fmt.Println(ans)
}
