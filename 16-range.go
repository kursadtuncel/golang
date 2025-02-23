package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

//range iterates over elements in a variety of built-in data structures.
//here we use range to sum the numbers in a slice. arrays work like this too.
//range on arrays and slices provides both the index and value for each entry.
//above we didn't need the index, so we ignoted it with the blank ideentifier _.
//Sometimes we actually want the indexes though.
//range on map iterates over key/value pairs.
//range can also iterate over just the keys of map.
//range on strings iterates over unicode code points. the first value is the starting byte-
//+ index of the rune and the second the rune itself.
