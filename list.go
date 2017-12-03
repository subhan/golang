package main

import (
	"fmt"
)

type il []int
type sl []string

type list interface {
	add_element()
	pop_element()
	size()
}

func main() {
	ml := il{0, 5, 10}
	kl := sl{}
	ml.add_element(15)
	kl.add_element(15)
	fmt.Println(ml, kl)
}

func (l il) add_element(ele int) {
	l = append(l, ele)
}

func (l sl) add_element(ele string) {
	l[len(l)] = ele
}
