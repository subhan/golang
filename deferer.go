package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("First")
    val := 45
    defer func() {
        if err := recover(); err != nil{
            fmt.Println("ho no")
        }
        fmt.Println("later", val)
    }()
    val = 22
    //panic("oops")
    kimaya(os.Args[1])
}

func kimaya(arg string) {
    fmt.Println(arg+ ": kit")
}
