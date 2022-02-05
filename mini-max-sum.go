package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
func miniMaxSum(arr []int32) {
	// Write your code here
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var sum1 = 0
	var sum2 = 0

	underArr1 := arr[:len(arr)-1]
	underArr2 := arr[1:]
	for _, v := range underArr1 {
		sum1 += int(v)
	}
	for _, i := range underArr2 {
		sum2 += int(i)
	}

	fmt.Println(sum1, sum2)
}

/*
Sample Input
1 2 3 4 5

Sample Output
10 14

*/
