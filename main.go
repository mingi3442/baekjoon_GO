package main

import (
	"fmt"
)

func main() {
	var total int
	var result int = 0
	fmt.Scanf("%d", &total)
	for i := 1; total+1 > i; i++ {
		result += i
	}
	fmt.Println(result)
}
