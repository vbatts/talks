// +build ignore

package main

func main() {
	// START1 OMIT
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	println(<-c)
	// STOP1 OMIT
}
