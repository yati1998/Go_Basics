package main

import "fmt"

func main(){
	i, j := 10,20
	p := &i
	fmt.Println("Original value of i=",*p)
	*p = i+10
	fmt.Println("New value of i= ", i)

	p = &j
	fmt.Println("original value of j=",j)
	*p = j+10
	fmt.Println("New value of j=", *p)
}
