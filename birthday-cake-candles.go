package main

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	m := map[int32]int32{}
	var maxCnt int32
	var freq int32
	var count int32
	for _, a := range candles {
		m[a]++
		if m[a] > maxCnt {
			maxCnt = m[a]
			freq = a
		}
		if a == freq {
			count++
		}

	}
	return count
}

/*

Sample Input
4
3 2 1 3

Sample Output
2

*/
