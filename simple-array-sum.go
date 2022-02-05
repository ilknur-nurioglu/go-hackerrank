package main

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var sum int32 = 0
	for _, v := range ar {
		sum += v
	}
	return sum
}
