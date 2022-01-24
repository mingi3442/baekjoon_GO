package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%d", &num2)
	fmt.Println(num1 * (num2 % 10))
	fmt.Println(num1 * ((num2 / 10) % 10))
	fmt.Println(num1 * (num2 / 100))
	fmt.Println(num1 * num2)
}
