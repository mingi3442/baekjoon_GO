package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var numA, numB, numC int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n%d\n%d\n", &numA, &numB, &numC)
	var result = numA * numB * numC
	//1~9 숫자를 위해 배열 선언
	var nums = make([]int, 10)
	for true {
		//result를 10으로 나눈 나머지랑 같은 값을 갖고 있는 index의 숫자 증가
		nums[result%10]++

		result = result / 10
		if result == 0 {
			break
		}
	}

	for i := 0; 10 > i; i++ {
		fmt.Println(nums[i])
	}
}
