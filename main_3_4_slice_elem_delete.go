package main

import "fmt"

func main() {
	a := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		a = append(a, i)
	}

	n := 50
	// a = append(a[:n], a[n+1:]...) // 添字n（50）の要素を削除
	a = a[:n+copy(a[n:], a[n+1:])] // 添字n（50）の要素を削除
	fmt.Println(a)
}
