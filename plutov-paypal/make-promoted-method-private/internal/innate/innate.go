package innate

import "fmt"

type Parent struct {
	Value int64
}

func (p *Parent) SomeMethod() {
	fmt.Println("Hi, details are:", p)
}
