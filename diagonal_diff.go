package main

import (
    "fmt"
    "os"
    "strings"
    "bufio"
    "strconv"
)

func main() {
    var N int
    fmt.Scanf("%d", &N)
    multi := make([][]int, 0, N)
    reader := bufio.NewReader(os.Stdin)
    for i := 0; i<N; i++ {
        text, _ := reader.ReadString('\n')
        text = strings.Replace(text, "\n", "", -1)
        temp :=make([]int, 0, N)
        for _, val := range strings.Split(text, " ") {
            val, _ := strconv.Atoi(val)
            temp = append(temp, val)
        }
        multi = append(multi, temp)
    }
    fmt.Println(diagonal_difference(multi))
}

func diagonal_difference(arr [][]int) int {
    d1 := 0
    d2 := 0
    for i, row := range arr{
        d1 += row[i]
        d2 += row[len(row)-i-1]
    }
    sum := d1-d2
    if sum < 0 {
        sum = -sum
    }
    return sum
}
