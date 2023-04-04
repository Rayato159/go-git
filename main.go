package main

import "fmt"

// 1 add files
// 2 commit
// 3 push

func sum(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("Hello Github")
	fmt.Println(sum(1, 2))
	fmt.Println(mul(1, 2))
}
