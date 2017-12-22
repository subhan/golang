package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "bufio"
)

func main() {

    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
    var array [2]int
    var n,m int
    for i, val := range strings.Split(text, " ") {
        val, _ := strconv.Atoi(val)
        array[i] = val
    }
    n, m = array[0], array[1]
    A := generate_array(n)
    B := generate_array(m)
    fmt.Println(A)
    fmt.Println(find_mininum_value(A))
    fmt.Println(B)
    fmt.Println(find_mininum_value(B))
}


func all_are_primes(array []int) bool {
    for _, val := range array {
        if is_prime(val) == false{
            return false
        }
    }
    return true
}


func gcd(array []int) int {
    small, big := array[0], array[1]
    if small > big {
        small, big = big, small
    }
    reminder = big % small
    if reminder == 0 {
        for i := 2; i < len(array); i++ {
        }
    }
}

func find_mininum_value(array []int) int {
    pr := prime_generator{value:2}
    primes := 1
    for all_are_primes(array) == false {
        flag := true
        primes *= pr.value
        for i, val:= range array {
            if val % pr.value == 0 {
                array[i] = int(val/pr.value)
                flag = false
            }
        }
        if flag {
            pr = pr.generator()
        }
    }
    for _, val := range array {
        primes *= val
    }
    return primes
}


func generate_array(size int) []int {
    array := make([]int, 0, size)

    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
    for _, val := range strings.Split(text, " ") {
        val, _ := strconv.Atoi(val)
        array = append(array, val)
    }
    return array
}
