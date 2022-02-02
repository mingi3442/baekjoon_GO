package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var total int
	var max int
	var sum float64
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &total)
	var scores = make([]int, total)
	//최고점 찾기
	for i := 0; total > i; i++ {
		fmt.Fscanf(reader, "%d ", &scores[i])
		if max < scores[i] {
			max = scores[i]
		}
	}
	//모든 점수를  => 점수/최고점 * 100 으로 변환 후 총 점수 sum에 더함
	for j := 0; total > j; j++ {
		sum += float64(scores[j]) / float64(max) * 100.0
	}
	//새로운 점수를 기준으로 평균
	fmt.Println(sum / float64(total))
}
