package main

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	var totalPoint [2]int32
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			totalPoint[0]++
		}
		if a[i] == b[i] {
			continue
		}
		if a[i] < b[i] {
			totalPoint[1]++
		}

	}
	return totalPoint[:]
}
