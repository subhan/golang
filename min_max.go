package main

import (
    "bufio"
    "os"
    "strings"
    "fmt"
    "strconv"
)

type ARR []int



func main() {
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
    var arr ARR
    for _, val := range strings.Split(text, " ") {
        val, _ := strconv.Atoi(val)
        arr = append(arr, val)
    }
    min_val := arr.min()
    max_val := arr.max()
    fmt.Println(max_val, min_val)
}

func (arr ARR) max() int {
    var temp int
    var flag bool

    for i :=0; i < len(arr); i++ {
        var val int
        for j := 0; j < len(arr); j++ {
            if i != j {
                val += arr[j]
            }
        }
        if flag == false {
            temp = val
            flag = true
        }else if temp > val {
            temp = val
        }
    }
    return temp
}

func (arr ARR) min() int {
    var temp int
    var flag bool

    for i :=0; i < len(arr); i++ {
        var val int
        for j := 0; j < len(arr); j++ {
            if i != j {
                val += arr[j]
            }
        }
        if flag == false {
            temp = val
            flag = true
        }else if temp < val {
            temp = val
        }
    }
    return temp
}
