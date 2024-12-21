package main

import (
	"fmt"
	"math/rand"
)

func randomGenerator(ch chan int) {
	for {
		ch <- rand.Int()
	}
}

func main() {
	ch := make(chan int)

	go randomGenerator(ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	close(ch)
}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func randomGenerator(ch chan int) {
// 	for {
// 		ch <- rand.Int()
// 	}
// }

// func main() {
// 	ch := make(chan int)

// 	go randomGenerator(ch)

// 	for i := 0; i < 10; i++ {
// 		select {
// 		case num := <-ch:
// 			fmt.Println(num)
// 		case <-time.After(1 * time.Second):
// 			fmt.Println("Таймаут")
// 			return
// 		}
// 	}

// 	close(ch)
// }
