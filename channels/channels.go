package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	// receiving goroutine
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	// sender goroutine
	go func() {
		i := 42
		ch <- i
		i = 27
		wg.Done()
	}()
	wg.Wait()

	fmt.Println()
	chanInLoopExample()
	fmt.Println()
	receiveOnlySendOnlyExample()
	fmt.Println()
	bufferedChanExample()
	fmt.Println()
	chanForRangeLoopExample()
	fmt.Println()
	chanOkExample()
	fmt.Println()
}

func chanInLoopExample() {
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch <- 42 // pauses the execution of the goroutine until there's space available in the channel
			wg.Done()
		}()
	}
	wg.Wait()
}

func receiveOnlySendOnlyExample() {
	ch := make(chan int)

	wg.Add(2)
	// receive only
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		// ch <- 27 // invalid operation: cannot send to receive-only channel ch (variable of type <-chan int)
		wg.Done()
	}(ch)
	// send only
	go func(ch chan<- int) {
		ch <- 42
		// fmt.Println(<-ch) // invalid operation: cannot receive from send-only channel ch (variable of type chan<- int)
		wg.Done()
	}(ch)

	wg.Wait()
}

func bufferedChanExample() {
	ch := make(chan int, 50) // channel has an internal data store (buffer) which can store 50 integers
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27 // if no buffer is given for the channel, then this goroutine is blocked here, and it can never complete -> deadlock error
		// with the buffer given, this message is lost, because there's only one receiver, which stops after reading the first integer
		wg.Done()
	}(ch)
	wg.Wait()
}

func chanForRangeLoopExample() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 77
		ch <- 78
		close(ch) // closed channels can't be reopened
		wg.Done()
	}(ch)
	wg.Wait()
}

func chanOkExample() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 77
		ch <- 78
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
