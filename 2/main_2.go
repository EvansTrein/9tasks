package main

import (
	"fmt"
	"math/rand"
)

func SliceExample(s []int) (evenSlice []int) {
	for _, el := range s {
		if el%2 == 0 {
			evenSlice = append(evenSlice, el)
		}
	}
	return evenSlice
}

func AddElements(s []int, num int) []int {
	newSlice := make([]int, len(s))  // IDE ругалась на newSlice := make([]int, len(s), len(s))
	copy(newSlice, s)
	newSlice = append(newSlice, num)
	return newSlice
}

func CopySlice(s []int) []int {
	newSlice := make([]int, len(s)) // IDE ругалась на newSlice := make([]int, len(s), len(s))
	copy(newSlice, s)

	return newSlice
}

func RemoveElement(s []int, indx int) ([]int, error) {
	if indx > len(s) {
		return nil, fmt.Errorf("индекс не может быть больше, чем кол-во элементов слайса")
	}

	newSlice := make([]int, 0, len(s))
	newSlice = append(newSlice, s[:indx]...)
	newSlice = append(newSlice, s[indx+1:]...)

	return newSlice, nil
}

func main() {
	// 1. Создайте слайс целых чисел originalSlice, содержащий 10 произвольных значений, которые генерируются случайным образом (при каждом запуске должны получаться новые значения)
	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}
	fmt.Println("originalSlice:", originalSlice)
	fmt.Println("<================================>")

	// 2. Напишите функцию sliceExample, которая принимает слайс и возвращает новый слайс, содержащий только четные числа из исходного слайса
	fmt.Println("только четные:", SliceExample(originalSlice))
	fmt.Println("<================================>")

	// 3. Напишите функцию addElements, которая принимает слайс и число. Функция должна добавлять это число в конец слайса и возвращать новый слайс.
	fmt.Println("добавить число в конец:", AddElements(originalSlice, 6666))
	fmt.Println("<================================>")

	// 4. Напишите функцию copySlice, которая принимает слайс и возвращает его копию. Убедитесь, что изменения в оригинальном слайсе не влияют на его копию.
	fmt.Printf("originalSlice: %v --> %p\n", originalSlice, originalSlice)
	sliceNew := CopySlice(originalSlice)
	fmt.Printf("sliceNew: %v --> %p\n", sliceNew, sliceNew)
	originalSlice[0] = 20000
	fmt.Println("originalSlice после изменения:", originalSlice)
	fmt.Println("sliceNew после изменения originalSlice:", sliceNew)
	fmt.Println("<================================>")

	// 5. Напишите функцию removeElement, которая принимает слайс и индекс элемента, который нужно удалить. Функция должна возвращать новый слайс без элемента по указанному индексу.
	fmt.Println("originalSlice:", originalSlice)
	fmt.Print("удаляем первый элемент:")
	fmt.Println(RemoveElement(originalSlice, 0))

	fmt.Print("удаляем последний элемент:")
	fmt.Println(RemoveElement(originalSlice, 9))

	fmt.Print("передаем неверный индекс:")
	fmt.Println(RemoveElement(originalSlice, 15))
}
