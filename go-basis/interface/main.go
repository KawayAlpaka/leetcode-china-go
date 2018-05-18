package main

import "fmt"

type iA interface {
	AF()
}

// Test ...
func Test(a iA) {
	a.AF()
}

// A ...
type A struct {
}

// AF ...
func (a *A) AF() {
	fmt.Println("A.AF")
}

func main() {
	fmt.Println(string(2) + "1")
	var i iA
	i = new(A)
	Test(i)
	var a A
	i = &a
	Test(i)
}
