package main

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */
func gradingStudents(grades []int32) []int32 {
	// Write your code here
	finalGrades := make([]int32, len(grades))
	for i := 0; i <= len(grades)-1; i++ {
		if grades[i] >= 38 && (grades[i]+2)%5 == 0 {
			finalGrades[i] = grades[i] + 2
		} else if grades[i] == 99 || (grades[i] > 38 && (grades[i]+1%5 == 0)) {
			finalGrades[i] = grades[i] + 1
		} else {
			finalGrades[i] = grades[i]
		}
	}
	return finalGrades
}

/*
Input
73
67
38
33

Output
75
67
40
33
*/
