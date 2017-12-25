package main

import (
    "bufio"
    "io"
    "log"
    "os"
    "fmt"
)

func main() {
    nBytes, nChunks := int64(0), int64(0)
    filename, _ := os.Open(os.Args[1])
    defer filename.Close()

    r := bufio.NewReader(filename)
    buf := make([]byte, 0, 4*1024)
    for {
        n, err := r.Read(buf[:cap(buf)])
        buf = buf[:n]
        if n == 0 {
            if err == nil {
                continue
            }
            if err == io.EOF {
                break
            }
            log.Fatal(err)
        }
        nChunks++
        nBytes += int64(len(buf))
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
    }
}


func process_data(buf []byte, count_data map[string]int)
