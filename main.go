package main

import "fmt"

// 1 add files
// 2 commit
// 3 push

func sub(a, b int) int {
	return a - b
}

func sum(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func dev(a, b float64) float64 {
	return a / b
}

func display(msg string) {
	fmt.Println(msg)
}

func main() {
	fmt.Println("Hello Github")
	fmt.Println(sum(1, 2))
	fmt.Println(sub(1, 2))
	fmt.Println(mul(1, 2))
	fmt.Println(dev(1, 2))
	display("Hello Gitflow")
}
