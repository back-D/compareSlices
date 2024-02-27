package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

var (
	slice1 = []string{"bar", "baz", "cos", "doz", "foo"}
	slice2 = []string{"dos", "foo", "gob"}
	str    = "bar"
	map1   = map[string]struct{}{
		"foo": {},
		"bar": {},
		"baz": {},
		"cos": {},
		"doz": {},
	}
)

func main() {
	sort.Strings(slice2)
	fmt.Println(slice2)
}

func SliceContains(s1 []string, s string) bool {
	if slices.Contains(s1, s) {
		return true
	}
	return false
}

func MapContains(s2 []string, m1 map[string]struct{}) bool {
	for _, s := range s2 {
		if _, found := m1[s]; found {
			return true
		}
	}

	return false
}

func ContainsSorted[T cmp.Ordered](haystack []T, needle T) bool {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

func IsIntersect[T cmp.Ordered](sliceA, sliceB []T) bool {
	var indexA, indexB int
	for indexA < len(sliceA) && indexB < len(sliceB) {
		if sliceA[indexA] == sliceB[indexB] {
			return true
		} else if sliceA[indexA] > sliceB[indexB] {
			indexB++
		} else if sliceA[indexA] < sliceB[indexB] {
			indexA++
		}
	}
	return false
}
