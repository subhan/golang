package main

import "fmt"

type Animal interface {
    Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
    return "bow bow"
}

type Cat struct {
}

func (c Cat) Speak() string {
    return  "Meow!"
}

type Lama struct {
}

func (l Lama) Speak() string {
    return "***!"
}

func main() {
    animals := []Animal{Cat{}, Dog{}, Lama{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
