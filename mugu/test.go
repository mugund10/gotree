package main

import "fmt"

func main() {
    ch := make(chan int)
    go rec(3, ch)
    
    // Read values from the channel until it's closed
    for val := range ch {
        fmt.Println(val)
    }
	
}

func rec(a int, ch chan<- int) {
    // Base case: Stop recursion after a certain condition
    // if a >= 6 {
    //     close(ch)  // Close channel to prevent deadlock
    //     return
    // }
    
    ch <- a  // Send current value to channel
    go rec(a+1, ch)  // Increment `a` and spawn new goroutine
}