package main

import (
	"container/heap"
	"fmt"
	"github.com/kaleabAlemayehu/kzip/huffman"
	"os"
)

func main() {
	hash_map := GenerateHashMap()
	huff_heap := huffman.NewHuffHeap(hash_map)
	heap.Init(huff_heap)
	for huff_heap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(huff_heap))
	}
}

func GenerateHashMap() map[byte]int {

	var hash_map map[byte]int = make(map[byte]int)
	byt, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	for _, b := range byt {
		i, ok := hash_map[b]
		if ok {
			hash_map[b] = i + 1
		} else {
			hash_map[b] = 1
		}
	}
	return hash_map
}
