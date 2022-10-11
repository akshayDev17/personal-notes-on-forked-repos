package internal

import (
	"fmt"
	"make-promoted-method-private/internal/innate"
)

type Child struct {
	parent  innate.Parent
	Details string
}

func (c *Child) HelloChild() {
	fmt.Println("Hi, I am inside a child, hence can call the get details method")
	c.parent.SomeMethod()
	fmt.Println("Exiting the child")
}

func NewChild(Details string) Child {
	return Child{
		parent: innate.Parent{
			Value: 37,
		},
		Details: Details,
	}
}
