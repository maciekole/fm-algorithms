package main

import (
	"math"
)

func CrystalBallsSearch(arr []bool) (int, error) {
	jump := int(math.Floor(math.Sqrt(float64(len(arr)))))
	i := jump
	for j := i; j < len(arr); j += jump {
		if arr[j] {
			i = j - jump
			break
		}
	}

	for k := 0; k <= jump && i < len(arr); k++ {
		if arr[i] {
			return i, nil
		}
		i++
	}
	return -1, nil
}
