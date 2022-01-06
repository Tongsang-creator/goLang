package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	//fmt.Println(reflect.TypeOf(input))
	intVar, _ := strconv.Atoi(input)
	//fmt.Println(intVar)
	//fmt.Println(intVar % 2)
	if intVar > 2 && (intVar%2 == 0) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
