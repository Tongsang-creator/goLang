package main

import (
	"fmt"
	"math"
)

func inputData() int {
	var num1, num2, num3 float64
	fmt.Scanf("%f %f %f\n", &num1, &num2, &num3)
	return int(math.Ceil(num1/num3) * math.Ceil(num2/num3))
}

func main() {
	total := inputData()
	if total <= 0 {
		fmt.Println(1)
	} else {
		fmt.Println(total)
	}
}
