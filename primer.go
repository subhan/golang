package main

import "math"



type prime_generator  struct {
    value int
    next *prime_generator
}


func (p *prime_generator) generator () prime_generator {
    i := p.value
    for {
        i++
        if is_prime(i) {
            p.next = &prime_generator{value:i}
            break
        }
    }
    return *p.next
}


func is_prime(num int) bool {
    if num <= 2 {
        return true
    }
    for i := 2; i<=int(math.Pow(float64(num), 0.5)); i++ {
         if  num % i == 0 {
            return false
        }
    }
    return true
}


