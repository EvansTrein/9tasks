package main

import (
	"fmt"
	"math/rand"
)

// Напишите генератор случайных чисел используя небуфферизированный канал
func RandNumsGenerator(count, diapason int) <-chan int {
	nums := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			nums <- rand.Intn(diapason)
		}
		close(nums)
	}()
	return nums
}

func main() {
	for num := range RandNumsGenerator(10, 100) {
		fmt.Println(num)
	}
}

// Напишите генератор случайных чисел используя небуфферизированный канал
// func RandNumsGenerator(ch chan int) {
// 	iterCount := 0
// 	for {
// 		iterCount++
// 		ch <- rand.Intn(1000)
// 		if iterCount == 20 {
// 			time.Sleep(time.Second * 2)
// 			close(ch)
// 		}
// 	}
// }

// func main() {
// 	ch := make(chan int)

// 	go RandNumsGenerator(ch)

// 	for {
// 		// select приоритет операций:
// 		// 1 - чтение из канала (если несколько case читают, то планировщик выбырет рандомный), unblock операция
// 		// 2 - запись в канал, block операция
// 		// 3 - default
// 		select {
// 		case num := <-ch:  // чтение из канала, unblock операция, на 20 итерации горутнна randomGenerator будет стоять 2 секунды
// 			fmt.Println(num)
// 		case <-time.After(time.Second * 1):  // в это время сработает этот case, так как тут 1 секунда
// 			fmt.Println("Таймаут")
// 			return
// 		}
// 	}
// }
