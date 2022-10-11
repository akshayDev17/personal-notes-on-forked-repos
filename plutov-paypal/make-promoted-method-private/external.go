package main

import (
	"make-promoted-method-private/internal"
)

func main() {
	var sampleChild internal.Child
	sampleChild = internal.NewChild("Cristiano Ronaldo")
	sampleChild.HelloChild()
	// sampleChild.someMethod() // compile issue
}
