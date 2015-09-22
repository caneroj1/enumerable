# enumerable
An enumerable package for go

# usage
```go
package main

import (
	"fmt"
	"github.com/caneroj1/enumerable"
)

func main() {
	// square each of the values in the slice and return a new slice
	myInts := []int{1, 2, 3, 4, 5}
	squared, err := enumerable.Map(myInts, func(idx, val int) int {
		return val * val
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(squared)
	// err is nil
	// squared = [1 4 9 16 25]

	// check if all of the words in the slice are at most 3 letters long
	myWords := []string{"cat", "dog", "man", "house"}
	result, err := enumerable.All(myWords, func(idx int, val string) bool {
		return len(val) <= 3
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// err is nil
	// result = false

	// check if there are any true values in the array
	myBools := []bool{true, false, false, false}
	result, err = enumerable.Some(myBools, func(idx int, val bool) bool {
		return val
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// err is nil
	// result = true

	// return a new slice containing the length of each word
	wordCounts, err := enumerable.Map(myWords, func(idx int, val string) int {
		return len(val)
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wordCounts)
	// err is nil
	// wordCount = [3 3 3 5]

	// outputs an error
	_, err = enumerable.Map(10, func(idx int, val string) int {
		return len(val)
	})
	if err != nil {
		fmt.Println(err)
	}
	// Enumerable Error: A slice needs to be the first parameter of Map.
}
```
