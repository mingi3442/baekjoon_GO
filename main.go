package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var total int
	var star string = "*"
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &total)
	for i := 0; total > i; i++ {
		fmt.Fprintf(writer, "%s\n", star)
		star += "*"
	}
	writer.Flush()
}
