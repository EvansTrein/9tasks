package main

import (
	"testing"
)

func TestNewCustomWaitGroup(t *testing.T) {
	var counterFact int
	wantCounter := 5
	
	t.Run("запускаем 5 задач", func(t *testing.T) {
		wg := NewCustomWaitGroup(5)

		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				counterFact += 1
			}()
		}

		wg.Wait()

		if counterFact != wantCounter {
			t.Errorf("количество выполеннных задач - %v не совпадает с ожидаемым - %v", counterFact, wantCounter)
		}
	})
}
