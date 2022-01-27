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
	for i := 1; total+1 > i; i++ {
		fmt.Fprintln(writer, i)
	}
	writer.Flush()
}
