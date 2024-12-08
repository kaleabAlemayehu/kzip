package main

import (
	"fmt"
	"os"
)

func main() {
	hash_map := GenerateHashMap()
	for key, value := range hash_map {
		fmt.Printf("Key: %s, value: %d\n", key, value)
	}
}

func GenerateHashMap() map[string]int {

	var hash_map map[string]int = make(map[string]int)
	byt, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	for _, b := range byt {
		i, ok := hash_map[string(b)]
		if ok {
			hash_map[string(b)] = i + 1
		} else {
			hash_map[string(b)] = 1
		}
	}
	return hash_map
}
