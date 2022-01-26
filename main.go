package main

import "fmt"

func main() {
	var hour, min int
	fmt.Scan(&hour, &min)
	if min >= 45 {
		fmt.Printf("%d %d\n", hour, min-45)
	} else {
		if hour == 0 {
			fmt.Printf("%d %d\n", 23, min+15)
		} else {
			fmt.Printf("%d %d\n", hour-1, min+15)
		}
	}
}
