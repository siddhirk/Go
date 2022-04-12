package main

import "fmt"

func main1() {
	var good string = "hello"
	var pointer *string = &good

	fmt.Println("good =", good)
	fmt.Println("pointer =", pointer)

	fmt.Println("*pointer =", *pointer)

	*pointer = "hii"
	fmt.Println("*pointer =", *pointer)

	fmt.Println("good =", good)
}
