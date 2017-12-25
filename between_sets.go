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
    min := find_mininum_value(A)
    max := find_gcd(B)
    if min > max {
        fmt.Println(0)
        os.Exit(1)
    }
    val := min
    count := 0
    for val >= min && val <= max {
        if (div_byA(A, val) == true) && (Bdiv_by_val(B, val) == true) {
            count ++
        }
        val ++
    }
    fmt.Println(count)
}

func div_byA(arr []int, d int) bool {
    for _, val := range arr {
        if d % val != 0{
            return false
        }
    }
    return true
}

func Bdiv_by_val(arr []int, d int) bool {
    for _, val := range arr {
        if val % d != 0{
            return false
        }
    }
    return true
}

func all_are_primes(array []int) bool {
    for _, val := range array {
        if is_prime(val) == false{
            return false
        }
    }
    return true
}


func find_gcd(arr []int) int {
    small, big := arr[0], arr[1]
    if small > big {
        small, big = big, small
    }
    div := gcd(small, big)
    for _, val := range arr[2:] {
        div = gcd(div, val)
    }
    return div
}

func gcd(small int, big int) int {
    reminder := big % small
    if reminder == 0 {
        return small
    } else {
        return gcd(reminder, small)
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
