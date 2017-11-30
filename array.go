package main

import "fmt"

func main() {
	scores := make([]int, 10)
	//scores = scores[0:8]
	//scores[1] = 9033
	scores = append(scores, 2334)
	fmt.Println(scores)
}
