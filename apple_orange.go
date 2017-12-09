package main


import (
    "fmt"
    "bufio"
    "strconv"
    "strings"
    "os"
    "io"
)


func main() {
    var s,t,a,b,m,n int
    arr := read_input()
    s, t = arr[0][0], arr[0][1]
    a, b = arr[1][0], arr[1][1]
    m, n = arr[2][0], arr[2][1]
    apples := arr[3]
    oranges := arr[4]
    if a > s || s > t || t > b{
        fmt.Println(0)
        fmt.Println(0)
    }
    process(s,t,a,b,m,n,apples,oranges)
}


func process(s int, t int, a int, b int, m int, n int, apples []int, oranges []int) {
    total := 1
    for _, val := range apples {
        val = a+val
        if  val >= s && val <= t {
            fmt.Println(total)
            total ++
        }
    }
    total = 1
    for _, val := range oranges {
        val = a+val
        if  val >= s && val <= t {
            fmt.Println(total)
            total ++
        }
    }
}


func read_input() [5][]int {
    var arr [5][]int
    i := 0
    snr := bufio.NewScanner(os.Stdin)
    for snr.Scan(){
        text := snr.Text()
        err := snr.Err()
        if len(text) == 0  || err == io.EOF{
            fmt.Println(text)
            break
        }
        text = strings.Replace(text, "\n", "", -1)
        var temp []int
        for _, val := range strings.Split(text, " ") {
            val, _ := strconv.Atoi(val)
            temp = append(temp, val)
        }
        arr[i] = temp
        i++
    }
    return arr
}
