package main

import (
	"fmt"

	lru "github.com/lucasmls/ecommerce/golang-generic-lru"
)

type Pet struct {
	Name string
}

func main() {
	cache, err := lru.New[Pet](2)
	if err != nil {
		panic(err)
	}

	evicted := cache.Add("dog", Pet{
		Name: "Pingo",
	})

	fmt.Println(evicted) // false

	evicted = cache.Add("bird", Pet{
		Name: "John",
	})

	fmt.Println(evicted) // false

	evicted = cache.Add("salamander", Pet{
		Name: "Bob",
	})
	fmt.Println(evicted) // true

	_, ok := cache.Get("dog")
	if !ok {
		fmt.Println("dog not found")
	}

	bird, ok := cache.Get("bird")
	if !ok {
		fmt.Println("bird not found")
	}

	fmt.Println(bird)
}
