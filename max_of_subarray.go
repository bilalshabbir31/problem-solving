package main

import (
	"fmt"
)

func find_max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {

	arr := []int{2, 10, 1, 8, 2, 9, 1}
	j := 0
	i := 0
	k := 3

	max := -1

	for j < len(arr) {
		max = find_max(max, arr[j])
		if (j - i + 1) < k {
			j++
		} else if (j - i + 1) == 3 {
			i++
			j++
		}
	}

	fmt.Println(max)

}
