package main


import (
    "fmt"
    "bufio"
    "strings"
    "os"
    "strconv"
)


func main() {

    A := make([]int, 3)
    B := make([]int, 3)

    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
    for i, val := range strings.Split(text, " "){
        val, _ := strconv.Atoi(val)
        A[i] = val
    }
    text, _ = reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
    for i, val := range strings.Split(text, " "){
        val, _ := strconv.Atoi(val)
        B[i] = val
    }
    compare(A, B)
}

func compare(A []int, B []int) {
    points := map[string]int{"A":0, "B":0}
    for i := 0; i < len(A); i++ {
        if A[i] > B[i] {
            points["A"] += 1
        }else if A[i] < B[i] {
            points["B"] += 1
        }
    }
    a1, _ := points["A"]
    b2, _ := points["B"]
    fmt.Println(a1, b2)
}
