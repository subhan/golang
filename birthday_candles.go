package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	candles := make(map[string]int)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	var values []int
	for _, key := range strings.Split(text, " ") {
		val, _ := candles[key]
		temp, _ := strconv.Atoi(key)
		values = append(values, temp)
		candles[key] = val + 1
	}
	fmt.Println(candles[strconv.Itoa(max(values))])
}

func max(arr []int) int {
	temp := arr[0]

	for _, val := range arr[1:] {
		if temp < val {
			temp = val
		}
	}
	return temp
}
