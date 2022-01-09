package main

import (
	"fmt"
	"strconv"
)

const (
	minNum    = 2
	modulaNum = 2
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	intVar, _ := strconv.Atoi(input)
	if intVar > minNum && (intVar%modulaNum == 0) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
