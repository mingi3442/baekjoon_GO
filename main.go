package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var total, flag int
	var minNum, maxNum int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &total, &flag)

	defer writer.Flush()

	var numbers = make([]int, total)
	for i := range numbers {
		fmt.Fscanf(reader, "%d ", &numbers[i])
	}
	minNum = numbers[0]
	maxNum = numbers[0]
	for j := range numbers {
		if numbers[j] > maxNum {
			maxNum = numbers[j]
		}
		if numbers[j] < minNum {
			minNum = numbers[j]
		}
	}

	fmt.Fprintf(writer, "%d %d\n", minNum, maxNum)
}
