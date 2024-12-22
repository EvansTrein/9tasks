package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_Uint8ToFloat64Cube(t *testing.T) {
	oneCh := make(chan uint8, 10)
	twoCh := make(chan float64, 10)
	wantResult := []float64{1, 8, 27, 64, 125}
	factResult := make([]float64, 0, 5)

	t.Run("5 чисел", func(t *testing.T) {
		go func() {
			for i := uint8(1); i <= 5; i++ {
				oneCh <- i
				time.Sleep(time.Second * 1)
			}
			close(oneCh)
			close(twoCh)
		}()

		go Uint8ToFloat64Cube(oneCh, twoCh)

		for num := range twoCh {
			factResult = append(factResult, num)
		}

		if !reflect.DeepEqual(wantResult, factResult) {
			t.Errorf("полученные значения - %v не соответствуют ожидаемым - %v", factResult, wantResult)
		}
	})
}
