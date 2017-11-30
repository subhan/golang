package main

import (
    "fmt"
    "os"
)

func main() {
    arr := []string{"a", "b", "c", "d"}

    /*for _, val := range os.Args[1:] {
        arr = append(arr, val)
    }*/
    for i := 1; i < len(os.Args); i++ {
        arr = append(arr, os.Args[i])
    }
    //arr = append(arr, "z")
    for i, value := range arr {
        fmt.Println(i, value)
    }
    fmt.Println(len(os.Args))
}
