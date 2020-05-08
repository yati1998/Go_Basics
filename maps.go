package main

import "fmt"

func main() {
	grades := make(map[string]int)
	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 67
	fmt.Println(grades)
	fmt.Println("Timmy's grade =",grades["Timmy"])
	delete(grades,"Timmy")
	fmt.Println(grades)
	for k,v := range grades{
		fmt.Println(k,":",v)
	}
}
