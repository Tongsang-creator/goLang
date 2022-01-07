package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	intVar, _ := strconv.Atoi(input)
	if intVar > 2 && (intVar%2 == 0) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
