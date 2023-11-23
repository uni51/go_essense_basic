package main

import "fmt"

type Value int

func (v *Value) Add(n Value) {
	*v += n // レシーバをデリファレンスして代入
}

func main() {
	v := Value(1)
	v.Add(2)       // 自身を書き換える
	fmt.Println(v) // 3が出力される
}
