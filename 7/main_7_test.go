package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	oneCh := make(chan int, 3)
	twoCh := make(chan int, 3)
	threeCh := make(chan int, 3)
	wantResult := []int{1, 11, 111, 2, 22, 222, 3, 33, 333}

	t.Run("сливаем 3 канала", func(t *testing.T) {
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
	
		resultNums := []int{}
		for resEl := range allNums {
			resultNums = append(resultNums, resEl)
		}
	
		sort.Ints(resultNums)
		sort.Ints(wantResult)
	
		if !reflect.DeepEqual(resultNums, wantResult) {
			t.Errorf("MergeChannels() = %v, want %v", resultNums, wantResult)
		}
	})
}
