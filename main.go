package main

import (
	"fmt"
	"unicode"
)

type Numeric interface {
	~int8 | int16 | int32 | int64 | float32 | float64
}

type Smallint int8

func doubler[T Numeric](value T) T {
	return value * 2
}

func filterPositive[T Numeric](items []T) []T {
	var filtered []T
	for _, v := range items {
		if v > 0 {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func filter[T any](items []T, fx func(T) bool) []T {
	var filtered []T
	for _, v := range items {
		if fx(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	strings := []string{"My", "name", "is", "Blessed", "Sibanda"}
	strings = filter(strings, func(s string) bool {
		return unicode.IsUpper(rune(s[0]))
	})
	fmt.Println(strings)

	// Filter a slice of integers
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	ints = filter(ints, func(i int) bool {
		return i%3 == 0
	})
	fmt.Println(ints)

	ints2 := []int8{-4, -3, -2, -1, 0, 1, 2, 3, 4}
	ints2 = filterPositive(ints2)
	fmt.Println(ints2)

	floats := []float32{-4.5, -3.5, -2.5, -1.5, 0, 0.5, 1.5, 2.5, 3.5, 4.5}
	floats = filterPositive(floats)
	fmt.Println(floats)

	var four Smallint = 4
	fmt.Println(doubler(four))
}
