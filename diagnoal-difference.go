package main

import (
	"math"
)

func DiagonalDifference(arr [][]int32) int32 {
	d1 := []int32{}
	d2 := []int32{}
	d1Sum := 0
	d2Sum := 0
	for x := range arr {
		for y := range arr {
			if x == y {
				d1 = append(d1, arr[x][y])
				d1Sum += int(arr[x][y])
			}

			if x == len(arr)-y-1 {
				d2 = append(d2, arr[x][y])
				d2Sum += int(arr[x][y])
			}
		}
	}

	return int32(math.Abs(float64(d1Sum) - float64(d2Sum)))
}
