package main

import (
    "fmt"
    "os"
    "bufio"
    "io"
)


func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }
    defer file.Close()

    io.Copy(os.Stdout, bufio.NewReader(file))
}
