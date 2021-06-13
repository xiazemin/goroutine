package main

import (
	"sync"

	"interface/map/util"
)

func main() {
	wg := sync.WaitGroup{}
	count := 8000
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()
			c := make(chan int, 1)
			//defer close(c)
			c <- util.Add(i, i+1)
			<-c
		}(i)
	}
	wg.Wait()
}
