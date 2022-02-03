package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	var numbers = make([]int, 10)
	var rest []int
	for i := range numbers {
		fmt.Fscanf(reader, "%d\n", &numbers[i])
		var temp = numbers[i] % 42
		rest = validate(rest, temp)
	}
	fmt.Println(len(rest))
}

func validate(temps []int, temp int) []int {
	for _, n := range temps {
		if temp == n {
			return temps
		}
	}
	return append(temps, temp)
}
