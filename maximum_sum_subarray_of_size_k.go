package main

import (
    "fmt"
)

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func main() {

    // find max sum of subarray
    k := 3
    // arr =  [2,5,1,8,2,9,1]
    arr := []int{2, 5, 1, 8, 2, 9, 1}
    j := 0
    i := 0
    var max int = -1
    sum := 0
    for j < len(arr) {
        sum = sum + arr[j]
        if (j - i + 1) < k {
            j = j + 1
        } else if (j - i + 1) == k {
            max = Max(max, sum)
            sum = sum - arr[i]
            i = i + 1
            j = j + 1
        }
    }
    fmt.Print(max)
}
