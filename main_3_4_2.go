package main

import "fmt"

type Fruit int
type Animal int

const (
	Apple Fruit = iota
	Orange
	Banana
)

const (
	Monkey Animal = iota
	Elephant
	Pig
)

func main() {
	var fruit Fruit = Banana
	fmt.Println(fruit) // 2

	// fruit = Elephant // ここがコンパイルエラーになる
	// fmt.Println(fruit)
}
