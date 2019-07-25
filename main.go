package main

import "fmt"

func main() {
	var input string
	fmt.Printf("enter your name: ")
	fmt.Scanln(&input)
	fmt.Printf("GETTING DONE,%s!", input)
}
