package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func strsAndInts(strs ...string) {
    katra :=  " "
    fmt.Print(strs, " ")
    for _, sr := range strs {
        katra = katra + sr
    }
    fmt.Println(katra)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)
    strsAndInts("A", "B", "C")

    nums := []int{1, 2, 3, 4}
    sum(nums...)
}

