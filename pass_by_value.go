package main

import "fmt"

type person struct {
    firstName string
    lastName string
    contactInfo
}

type contactInfo struct {
    email string
    zipcode int
}

func main() {
    p := person{
        firstName: "mahabu",
        lastName: "subhani",
        contactInfo: contactInfo{email: "subhan@subhan.com", zipcode:560032},
    }
    fmt.Println(p)
    p.updateName("mahaboob","subhan")
    fmt.Println(p)
}


func (p *person) updateName(newFirstName string, newLastName string) {
    p.lastName = newLastName
    p.firstName = newFirstName
}
