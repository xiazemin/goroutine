package main

import (
	"sync"

	"interface/main/util"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go func(i int) {
			defer wg.Done()
			util.Add(i, i+1)
		}(i)
	}
	wg.Wait()
}
