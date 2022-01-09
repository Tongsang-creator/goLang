package main

import (
	"fmt"
	"strconv"
)

const (
	initData int = 0
	minusNum int = 2
)

type varInput struct {
	nLines int
	words  string
	err    error
	lenght int
}

func inputSys() string {
	var input string
	fmt.Scanf("%s\n", &input)
	return input
}

func isTooLong(words string) bool {
	if len(words) > 10 {
		return true
	}
	return false
}

func main() {
	var data1 varInput
	data1.nLines, data1.err = strconv.Atoi(inputSys())

	for i := initData; i < data1.nLines; i++ {
		data1.words = inputSys()
		data1.lenght = len(data1.words)
		if isTooLong(data1.words) {
			lenghtBetween := strconv.Itoa(data1.lenght - minusNum)
			fmt.Printf("%c%v%c\n", data1.words[0], lenghtBetween, data1.words[data1.lenght-1])
		} else {
			fmt.Println(data1.words)
		}

		//fmt.Println(data1)

	}

}
