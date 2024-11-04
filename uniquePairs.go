package main

import (
	"flag"
	"fmt"
	"sort"
)

func main() {

	iMin := flag.Int("iMin", 1, "i min")
	iMax := flag.Int("iMax", 9, "i max")
	jMin := flag.Int("jMin", 1, "j min")
	jMax := flag.Int("jMax", 9, "j max")

	flag.Parse()

	pairs := make(map[[2]int]bool)

	for i := *iMin; i <= *iMax; i++ {
		for j := *jMin; j <= *jMax; j++ {
			if i != j {
				pair := [2]int{i, j}

				sort.Ints(pair[:])

				// Add pair to map if it's not already present
				pairs[pair] = true
			}
		}
	}

	// Convert map keys to a slice for sorting
	var pairSlice [][2]int
	for pair := range pairs {
		pairSlice = append(pairSlice, pair)
	}

	// Sort the pairs by first element, then by second element
	sort.Slice(pairSlice, func(i, j int) bool {
		if pairSlice[i][0] == pairSlice[j][0] {
			return pairSlice[i][1] < pairSlice[j][1]
		}
		return pairSlice[i][0] < pairSlice[j][0]
	})

	// Print the sorted pairs
	fmt.Println("Unique pairs:")
	count := 0
	for _, pair := range pairSlice {
		fmt.Println(pair)
		count++
	}
	fmt.Println("Unique pairs count:", count)

}
