package main

import "fmt"

// range:
// - used for iterating over data structures like arrays, slices, maps, or strings
// - returns both index and value for arrays/slices
// - returns key and value for maps
// - returns index and rune for strings

func main() {
	numSlice := []int{2, 42, 532, 532, 53}

	// without range
	fmt.Println("------------")
	for i := 0; i < len(numSlice); i++ {
		fmt.Println(numSlice[i])
	}

	// using range - index only
	fmt.Println("------------")
	for i := range numSlice {
		fmt.Println(numSlice[i])
	}

	// using range - value only
	fmt.Println("------------")
	for _, val := range numSlice {
		fmt.Println(val)
	}

	// using range - both index and value
	fmt.Println("------------")
	for i, val := range numSlice {
		fmt.Printf("value at index %d is: %d\n", i, val)
	}

	fmt.Println("------------")
	// range with map
	mapEg := map[string]interface{}{
		"name": "Tejas",
		"age":  23,
		"add":  "Barshi",
	}

	fmt.Println("------------")
	for key, val := range mapEg {
		fmt.Printf("{%s : %v}\n", key, val)
	}

	fmt.Println("------------")
	// range on string - prints Unicode code points
	str2 := "amar thonge"
	for i, ch := range str2 {
		fmt.Printf("index: %d, character: %c, unicode: %d\n", i, ch, ch)
	}
	/*

	index: 0, character: a, unicode: 97
	index: 1, character: m, unicode: 109
	index: 2, character: a, unicode: 97
	index: 3, character: r, unicode: 114
	index: 4, character:  , unicode: 32
	index: 5, character: t, unicode: 116
	index: 6, character: h, unicode: 104
	index: 7, character: o, unicode: 111
	index: 8, character: n, unicode: 110
	index: 9, character: g, unicode: 103
	index: 10, character: e, unicode: 101

	*/
}
