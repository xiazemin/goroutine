package util

func Add(a, b int) chan int {
	for i := 0; i < 100000; i++ {

	}
	c := make(chan int)
	c <- a + b
	return c
}
