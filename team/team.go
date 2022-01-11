package main

import (
	"fmt"
	"strconv"
)

const (
	initData int = 0
	cutoff   int = 2
)

type varInput struct {
	nLines int
	err    error
	total  int
}

func inputSys() string {
	var input string
	fmt.Scanf("%s\n", &input)
	return input
}

func inputData() int {
	var num1, num2, num3 int
	fmt.Scanf("%d %d %d\n", &num1, &num2, &num3)
	return num1 + num2 + num3
}

func main() {
	var data1 varInput
	data1.nLines, data1.err = strconv.Atoi(inputSys())
	for i := initData; i < data1.nLines; i++ {
		if inputData() >= cutoff {
			data1.total++
		}
	}
	fmt.Println(data1.total)
}
