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
	var temp int = total
	for i := 0; total > i; i++ {
		fmt.Fprintln(writer, temp)
		temp -= 1
	}
	writer.Flush()
}
