package main

import "testing"

func Test_RandNumsGenerator(t *testing.T) {
	var nums []int
	count := 20
	diapason := 100

	t.Run("генерация 20 чисел", func(t *testing.T) {
		ch := RandNumsGenerator(count, diapason)

		for v := range ch {
			if v < 0 || v >= diapason {
				t.Errorf("сгенерированы числа выходящие за диапазон")
			}
			nums = append(nums, v)
		}

		if len(nums) != count {
			t.Errorf("сгенерировано %d, а должно быть %d", len(nums), count)
		}
	})
}
