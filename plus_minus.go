package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
)

func main() {
    var N int
    var p,n,z,t float32
    fmt.Scanf("%d", &N)
    arr := make([]int, 0, N)
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
    for _, val := range strings.Split(text, " ") {
        val, _ := strconv.Atoi(val)
        arr = append(arr, val)
        if val > 0 {
            p += 1
        }else if val < 0 {
            n += 1
        }else{
            z += 1
        }
        t++
    }
    fmt.Println(p/t)
    fmt.Println(n/t)
    fmt.Println(z/t)
}
