package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	//fmt.Scanf("%d", &x)
	//fmt.Scanf("%d", &y)
	if y-45 >= 0 {
		fmt.Printf("%d %d\n", x, y-45)
	} else if x == 0 {
		fmt.Printf("%d %d\n", 23, 15+y)
	} else {
		fmt.Printf("%d %d\n", x-1, 15+y)
	}
}
