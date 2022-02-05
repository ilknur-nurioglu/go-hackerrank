package main

import (
	"fmt"
	"strconv"
)

/* Note: - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
- 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock. */

func timeConversion(s string) string {
	// Write your code here
	format := s[len(s)-2:]
	strHour := s[:2]
	hour, err := strconv.Atoi(strHour)
	if err != nil {
		fmt.Println(err)
	}
	if format == "AM" && hour == 12 {
		return "00" + s[2:len(s)-2]
	}
	if format == "PM" && hour == 12 {
		return "12" + s[2:len(s)-2]
	}
	if format == "PM" {
		hour = hour + 12
	}
	return fmt.Sprintf("%02d", hour) + s[2:len(s)-2]

}

/*
Sample Input
07:05:45PM

Sample Output
19:05:45
*/
