package main

import (
    "fmt"
    "math/rand"
    "time"
)

func quicksort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }

    left, right := 0, len(arr)-1

    pivotIndex := rand.Int() % len(arr)

    // Move the pivot to the right
    arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

    // Move smaller elements to the left
    for i := range arr {
        if arr[i] < arr[right] {
            arr[i], arr[left] = arr[left], arr[i]
            left++
        }
    }

    arr[left], arr[right] = arr[right], arr[left]

    // Go crazy
    quicksort(arr[:left])
    quicksort(arr[left+1:])

    return arr
}

func main() {
    rand.Seed(time.Now().UnixNano())
    arr := []int{38, 4, 54, 104, 74, 14, 1, 94}
    fmt.Println("Original array:", arr)
    sortedArray := quicksort(arr)
    fmt.Println("Sorted array:  ", sortedArray)
}
