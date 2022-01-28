package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var total int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &total)
	for i := 0; total > i; i++ {
		for j := total - i - 1; j > 0; j-- {
			fmt.Fprintf(writer, " ")
		}
		for k := 0; i+1 > k; k++ {
			fmt.Fprintf(writer, "*")
		}
		fmt.Fprintf(writer, "\n")
	}
	writer.Flush()
}
