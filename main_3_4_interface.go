package main

import "fmt"

type Speaker interface {
	Speak() error
}

type Dog struct{}

func (d Dog) Speak() error {
	fmt.Println("BowWow")
	return nil
}

type Cat struct{}

func (c Cat) Speak() error {
	fmt.Println("Meow")
	return nil
}

func DoSpaeak(s Speaker) error {
	return s.Speak()
}

func main() {
	dog := Dog{}
	DoSpaeak(&dog) // BowWowが出力される

	cat := Cat{}
	DoSpaeak(&cat) // Meowが出力される
}
