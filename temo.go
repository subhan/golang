package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func main() {
    var fields string
    for snr.Scan(){
        line := snr.Text()
        err := snr.Err()
        fields += line
        if len(line) == 0  || err == io.EOF{
            break
        }
    }
    fmt.Printf("Fields: %q\n", fields)
}
