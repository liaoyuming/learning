package main

import "fmt"

//API is interface
type Person interface {
	Name() string
}

func Factory(t int) Person {
	if t == 1 {
		return &APerson{}
	} else if t == 2 {
		return &BPerson{}
	}
	return nil
}

type APerson struct{}

func (*APerson) Name() string {
	return fmt.Sprintf("I'm A")
}

type BPerson struct{}

func (*BPerson) Name() string {
	return fmt.Sprintf("I'm B")
}

func main() {
	factory := Factory(1)
	println(factory.Name())
	factory = Factory(2)
	println(factory.Name())
}