package main

import (
	"fmt"
)

const (
	Male     string = "IGNORE HIM!"
	FeMale   string = "CHAT WITH HER!"
	initData int    = 0
)

func inputSys() string {
	var input string
	fmt.Scanf("%s\n", &input)
	return input
}

func detect(input string) bool {
	var dict [30]int
	var count = initData
	for i := initData; i < len(input); i++ {
		if dict[input[i]-'a'] == 0 {
			count++
			dict[input[i]-'a']++
		}
	}

	if count%2 == 0 {
		return true
	} else {
		return false
	}

}

func main() {
	var input string

	input = inputSys()

	if detect(input) {
		fmt.Println(FeMale)
	} else {
		fmt.Println(Male)
	}
}
