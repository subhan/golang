package main

import (
    "fmt"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)
    for i :=0; i < N; i++ {
        var s string
        for j := i+1; j < N; j++{
            s += " "
        }
        for k :=N-1-i; k<N; k++ {
            s += "#"
        }
        fmt.Println(s)
    }
}
