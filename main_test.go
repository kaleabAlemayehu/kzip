package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGenerateHash(t *testing.T) {
	byt := read_file()

	hash_map := GenerateHashMap()
	var tests = []struct {
		input    []byte
		key      string
		expected int
	}{{byt, "X", 333}, {byt, "t", 223000}}
	for _, tt := range tests {
		test_name := fmt.Sprintf("%s\n", tt.key)
		t.Run(test_name, func(t *testing.T) {
			ans := hash_map[tt.key]
			if ans != tt.expected {
				t.Errorf("got %d, went %d\n", ans, tt.expected)
			}
		})
	}

}

func read_file() []byte {
	byt, err := os.ReadFile("./test.txt")
	if err != nil {
		panic(err)
	}
	return byt
}
