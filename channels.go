package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(somevalue int, c chan int){
	defer wg.Done()
	c <- somevalue*5
}
func main(){
	fooval := make(chan int,10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(i, fooval)
	}
	wg.Wait()
	close(fooval)

	for item := range fooval{
		fmt.Println(item)
	}
}
