package main

import "fmt"

type Mytest struct{
	x int
	y int 
}
func main(){
	v := Mytest{1,2}
	p := &v
	fmt.Println(v)
	p.x = 10
	p.y = 20
	fmt.Println(v)
}

