package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4}
    fmt.Println(arr)
    update(arr)
    fmt.Println(arr)
}

func update(arr []int) {
    arr[0] = 10
    arr[1] = 12
}
