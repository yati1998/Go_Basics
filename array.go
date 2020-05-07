package main

import "fmt"

func main(){
	var a []int
	var i int
	for i=0;i<5;i++ {
		a=append(a,i+1)
		fmt.Println(a)
	}
	sum := 0
	for i=0;i<5;i++ {
		sum =sum +a[i]
		fmt.Println(sum)
	}
	fmt.Println("sum =", sum)
}
