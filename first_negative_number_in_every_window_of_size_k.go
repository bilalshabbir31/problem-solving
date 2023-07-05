package main

import (
	"container/list"
)

func main() {
	arr := []int{1, -1, -10, 2, 5, -9, -2, -9, 10}

	i := 0
	j := 0

	queue := list.New()

	for j < len(arr) {
		if arr[j] < 0 {
			queue.PushBack(arr[j])
		}
		if (j - i + 1) < 3 {
			j = j + 1
		} else if (j - i + 1) == 3 {

			if (queue.Len()) == 0 {
				println(0)
			} else {
				println(queue.Front().Value.(int))
				if arr[i] == queue.Front().Value.(int) {
					queue.Remove(queue.Front())
				}

				i = i + 1
				j = j + 1
			}
		}
	}
}
