package main
import "fmt"

type boto interface {
   getPrint()
   printer()
}
type decks map[string]string
type numbers map[string]int

/*func main() {
    ml := decks{"firstname": "subhan", "lastname": "mahaboob",}
    kl := numbers{"one": 1, "two": 2, "three":3,}
    printer(ml)
    printer(kl)
}*/

func printer(b boto) {
    b.getPrint()
}

func (d decks) getPrint() {
    for key, value := range d {
        fmt.Println(key, ": ", value)
    }
}
func (n numbers) getPrint() {
    for key, value := range n {
        fmt.Println(key, ": ", value)
    }
}
