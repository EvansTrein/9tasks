package main

import (
	"fmt"
	"sync"
)

// Напишите программу на Go, которая сливает N каналов в один
func MergeChannels(channels ...<-chan int) <-chan int {
	wg := &sync.WaitGroup{}
	merged := make(chan int)

	go func() {
		
		for _, ch := range channels {
			wg.Add(1)
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for el := range ch {
					merged <- el
				}
			}(ch, wg)
		}

		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	oneCh := make(chan int, 3)
	twoCh := make(chan int, 3)
	threeCh := make(chan int, 3)

	for _, el := range []int{1, 11, 111} {
		oneCh <- el
	}
	close(oneCh)

	for _, el := range []int{2, 22, 222} {
		twoCh <- el
	}
	close(twoCh)

	for _, el := range []int{3, 33, 333} {
		threeCh <- el
	}
	close(threeCh)

	allNums := MergeChannels(oneCh, twoCh, threeCh)

	for el := range allNums {
		fmt.Println(el)
	}
}
