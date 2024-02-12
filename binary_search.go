package main

import (
	"math"
	"slices"
)

func BinarySearch(arr []int, v int) (bool, error) {
	// binary search works only on sorted arrays!
	sorted := slices.IsSorted(arr)
	if !sorted {
		slices.Sort(arr)
	}

	start := 0
	end := len(arr)

	for start < end {
		mid := int(math.Floor(float64(start) + (float64(end)-float64(start))/2))
		lookup := arr[mid]
		if lookup == v {
			return true, nil
		}
		if lookup > v {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return false, nil
}
