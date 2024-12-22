package main

import (
	"fmt"
	"time"
)

// Сделать кастомную waitGroup на семафоре, не используя sync.WaitGroup
type CustomWaitGroup struct {
	sem   chan struct{}
	count int
}

func NewCustomWaitGroup(count int) *CustomWaitGroup {
	return &CustomWaitGroup{
		sem:   make(chan struct{}, count),
		count: count,
	}
}

func (wg *CustomWaitGroup) Add(count int) {
	wg.count += count
	wg.sem <- struct{}{}
}

func (wg *CustomWaitGroup) Done() {
	<-wg.sem
	wg.count--
	if wg.count == 0 {
		close(wg.sem)
	}
}

func (wg *CustomWaitGroup) Wait() {
	for i := 0; i < wg.count; i++ {
		wg.sem <- struct{}{}
	}
}

func main() {
	wg := NewCustomWaitGroup(5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i, "started")
			time.Sleep(time.Second * 2)
			fmt.Println(i, "finished")
		}(i)
	}

	wg.Wait()
	fmt.Println("All finished")
}
