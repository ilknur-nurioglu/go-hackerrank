package main

import "fmt"

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	// Write your code here
	for i := 0; i < int(n); i++ {
		for j := 0; j < int(n)-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

/*
Sample Input
6

Sample Output

     #
    ##
   ###
  ####
 #####
######

*/
