package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"John": 21,
		"Bob":  18,
		"Mark": 33,
	}

	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("key: %v, value: %v\n", k, m[k])
	}
}
