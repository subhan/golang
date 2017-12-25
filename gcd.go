package main

import "fmt"

func main() {
    fmt.Println(gcd(14, 24))
}

func gcd(small int, big int) int {
    reminder := big % small
    if reminder == 0 {
        return small
    } else {
        return gcd(reminder, small)
    }
}
