package main

import (
	"testing"
)

func BenchmarkContainsSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ContainsSorted(slice1, str)
	}
}

func BenchmarkIsIntersect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsIntersect(slice1, slice2)
	}
}

func BenchmarkMapContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MapContains(slice2, map1)
	}
}

func BenchmarkSliceContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SliceContains(slice1, str)
	}
}
