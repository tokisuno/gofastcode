package main

import "fmt"

func main(){
    var intSlice = []int{1, 2, 3}
    fmt.Println(sumSlice[int](intSlice))

    var float32Slice = []float32{1, 2, 3}
    fmt.Println(sumSlice[float32](float32Slice))

    var emptySlice = []int{}
    fmt.Println(isEmpty(emptySlice))
}

func sumSlice[T int | float32 | float64](slice []T) T{
    var sum T
    for _, v:= range slice{
        sum += v
    }
    return sum
}

// We can infer the type of our generic parameter since we're checking if
// the size of the slice is 0
func isEmpty[T any](slice []T) bool {
    return len(slice)==0
}
