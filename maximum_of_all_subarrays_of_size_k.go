package main

import (
	"container/list"
	"fmt"
)

func main() {

	arr := []int{1, 3, -1, -3, -2,5, 3, 6, 7}
	j := 0
	i := 0
	k := 3
	queue := list.New()
	var sub_arr_max []int

	for j < len(arr) {

		for queue.Len() > 0 && queue.Back().Value.(int) < arr[j] {
			queue.Remove(queue.Back())
		}

		queue.PushBack(arr[j])

		if (j - i + 1) < k {
			j++
		} else if (j - i + 1) == 3 {
			sub_arr_max = append(sub_arr_max, queue.Front().Value.(int))
			if arr[i] == queue.Front().Value.(int) {
				queue.Remove(queue.Front())
			}
			i++
			j++
		}
	}

	fmt.Println(sub_arr_max)

}
