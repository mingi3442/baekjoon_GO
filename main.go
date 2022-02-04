package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var total int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &total)
	var arr = make([]string, total) // 입력을 받을 배열 선언
	var temp = make([]int, total)   // 점수를 넣을 배열 선언
	for i := 0; total > i; i++ {
		fmt.Fscanln(reader, &arr[i])
		var correct, score int // O의 유무와 점수를 위한 변수 선언

		for j := 0; len(arr[i]) > j; j++ {
			if fmt.Sprintf("%c", arr[i][j]) == "O" {
				correct++
				score += correct
			} else {
				correct = 0
			}
		}
		temp[i] = score
	}
	// fmt.Println(temp)
	for _, e := range temp {
		fmt.Println(e)
	}
}
