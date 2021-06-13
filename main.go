package main

import (
	"sync"

	"interface/util"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go func(i int) {
			util.Add(i, i+1)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
