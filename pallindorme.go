package main

import "fmt"

func main(){
	var num,temp,d int
	var rev int = 0
	fmt.Scanln(&num)
	temp = num
	for num > 0{
		d = num%10
		rev = rev*10+d
		num = num/10
	}
	if (rev == temp){
		fmt.Printf("Ïts a pallindorme number")
	}else{
		fmt.Printf("Its not a pallindormic number")
	}
}
