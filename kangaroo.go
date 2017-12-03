package main

import  (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

func main() {

    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text  = strings.Replace(text, "\n", "", -1)
    var arr [4]int
    for i, val := range strings.Split(text, " ") {
        val, _ := strconv.Atoi(val)
        arr[i] = val
    }
    x1 := arr[0]
    v1 := arr[1]
    x2 := arr[2]
    v2 := arr[3]
    check(x1, v1, x2, v2)
}

func check(x1 int, v1 int, x2 int, v2 int) {
    if x2 > x1 && v2 > v1 {
        fmt.Println("NO")
    } else {
        val1 := x1+v1
        val2 := x2+v2
        for val1 != val2 {
            val1 += v1
            val2 += v2
            if val1 > 10000 || val2 > 10000 {
                fmt.Println("NO")
                break
            }
        }
        if val1 == val2 {
            fmt.Println("YES")
        }
    }
}


