package main

import (
	"fmt"
	"time"
)

// Сделать конвейер чисел
// Даны два канала.
// В первый пишутся числа типа uint8. Нужно, чтобы числа читались из первого канала по мере поступления,
// затем эти числа должны преобразовываться в float64 и возводиться в куб и результат записывался во второй канал.


func Uint8ToFloat64Cube(inChan chan uint8, outChan chan float64) {
	for num := range inChan {
		cubed := float64(num) * float64(num) * float64(num)
		outChan <- cubed
	}
}

func main() {
	oneCh := make(chan uint8)
	twoCh := make(chan float64)

	go func() {
		for i := uint8(1); i <= 10; i++ {
			oneCh <- i
			time.Sleep(time.Second * 1)  // за это время функция успевает сработать и канал twoCh освобождается
		}
		close(oneCh)
		close(twoCh)
	}()

	go Uint8ToFloat64Cube(oneCh, twoCh)

	for num := range twoCh {
		fmt.Println(num)
	}
}
