package main

import "fmt"

func add(a int,b int) int{
	var sum int
	sum =a+b
	return sum
}
func sub(a int,b int) int{
	return a-b
}
func main() {
	result := add(5,10)
	fmt.Println("sum =", result)
	result = sub(10,5)
	fmt.Println("Difference =", result)
}
