package main

import "fmt"

func main() {
	fmt.Println("Hello")
	printFoo()
}

func printFoo() {
	fmt.Println("Process in foo")
}
