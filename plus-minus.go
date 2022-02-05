package main

import "fmt"

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
func plusMinus(arr []int32) {
	a := len(arr)

	positiveCount := 0
	for _, pn := range arr {
		if pn > 0 {
			positiveCount++
		}
	}
	negativeCount := 0
	for _, pn := range arr {
		if pn < 0 {
			negativeCount++
		}
	}
	zeroCount := 0
	for _, pn := range arr {
		if pn == 0 {
			zeroCount++
		}
	}
	var p float64 = float64(positiveCount) / float64(a)
	var n float64 = float64(negativeCount) / float64(a)
	var z float64 = float64(zeroCount) / float64(a)

	fmt.Printf("%.6f", p)
	fmt.Println()
	fmt.Printf("%.6f", n)
	fmt.Println()
	fmt.Printf("%.6f", z)

}
