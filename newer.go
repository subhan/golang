package main

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	ml := NewSaiyan("subhan", 2)
	println("name:", ml.Name)
	println("power:", ml.Power)
}

/*func NewSaiyan(name string, power int) *Saiyan {
    return &Saiyan{
        Name: name,
        Power: power,
    }
}
func (name string, power int) NewSaiyan() *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}*/

func NewSaiyan(name string, power int) Saiyan {
    return Saiyan{
        Name: name,
        Power: power,
    }
}
