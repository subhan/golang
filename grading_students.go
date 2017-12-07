package main

import (
    "fmt"
    /*"bufio"
    "strconv"
    "strings"
    "os"*/
)


func main() {
    var N, temp int
    fmt.Scanf("%d", &N)
    arr := make([]int, 0, N)
    for i := 0; i < N; i++ {
        fmt.Scanf("%d", &temp)
        arr = append(arr, temp)
    }
    grade(arr)
}

func grade(arr []int) {
    for _, val := range arr {
        if val < 38 {
            fmt.Println(val)
        }else {
            roundof := nextMultipleOf5(val)
            if roundof - val < 3 {
                fmt.Println(roundof)
            }else{
                fmt.Println(val)
            }
        }

    }
}


func  nextMultipleOf5(val int) int {
    return val + 5 - (val % 5)
}
