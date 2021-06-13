package main

import (
	"sync"

	"interface/chan/util"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go func(i int) {
			defer wg.Done()
			c := util.Add(i, i+1)
			<-c
		}(i)
	}
	wg.Wait()
}
