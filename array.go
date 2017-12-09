package main

import "fmt"

func main() {
	//scores := make([]int, 10)
	//scores = scores[0:8]
	//scores[1] = 9033
	//scores = append(scores, 2334)
    scores := [4]int{1,2,3,4}
	fmt.Println(scores)
    update(&scores)
	fmt.Println(scores)
}

func update(arr *[4]int) {
    arr[0] = 10
    arr[1] = 12
}
