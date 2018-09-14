package main

// START OMIT
var ngoroutine = 100000

func f(left, right chan int) { left <- 1 + <-right }

func main() {
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0      // bang!
	x := <-leftmost // wait for completion
	println(x)      // 100000
}

// STOP OMIT
