package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var total, sum int
	var str string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &total)
	fmt.Fscanf(reader, "%s", &str)

	nums := strings.Split(str, "")
	for i := 0; total > i; i++ {
		num, _ := strconv.Atoi(nums[i])
		sum += num
	}
	fmt.Println(sum)
}
