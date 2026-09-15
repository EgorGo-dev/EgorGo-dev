package main 

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 4; i++ {
			ch <-i
		}
		close(ch)
	}()
	go func() {
		defer wg.Done()
		for r := range ch {
		fmt.Println(r)
	}
	}()
	wg.Wait()
}