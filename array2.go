package main

import (
    "fmt"
    "os"
)

func main() {
    arr := make([]int, 10)
    //arr = append(arr, 4)
    //arr =  arr[0:8]
    arr[2] = 5
    fmt.Println(arr)
    fmt.Println(os.Args[1:])
}
