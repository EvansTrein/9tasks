package main

import "fmt"

// Реализуйте структуру данных StringIntMap, которая будет использоваться для хранения пар "строка - целое число".
type StringIntMap struct {
	data map[string]int
}

// создание структуры, нужно для инициализации мапы
func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

// 1. Добавление элемента: Метод Add(key string, value int), который добавляет новую пару "ключ-значение" в карту.
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

// 2. Удаление элемента: Метод Remove(key string), который удаляет элемент по ключу из карты.
func (m *StringIntMap) Remove(key string) error {
	if _, ok := m.data[key]; !ok {
		return fmt.Errorf("нет такого ключа")
	}

	delete(m.data, key)

	return nil
}

// 3. Копирование карты: Метод Copy() map[string]int, который возвращает новую карту, содержащую все элементы текущей карты.
func (m *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int, len(m.data))

	for k, v := range m.data {
		newMap[k] = v
	}

	return newMap
}

// 4. Проверка наличия ключа: Метод Exists(key string) bool, который проверяет, существует ли ключ в карте.
func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]

	return ok
}

// 5. Получение значения: Метод Get(key string) (int, bool), который возвращает значение по ключу и булевый флаг, указывающий на успешность операции.
func (m *StringIntMap) Get(key string) (int, bool) {
	value, ok := m.data[key]

	return value, ok
}

func main() {
	newStringIntMap := NewStringIntMap()
	fmt.Println(newStringIntMap)

	newStringIntMap.Add("one", 1)
	newStringIntMap.Add("two", 2)
	newStringIntMap.Add("three", 3)
	fmt.Println(newStringIntMap)
	fmt.Println("<================================>")

	newStringIntMap.Remove("one")
	newStringIntMap.Remove("xxxxxxxx")
	fmt.Println(newStringIntMap)
	fmt.Println("<================================>")

	fmt.Printf("%v -> %p\n", newStringIntMap.data, newStringIntMap.data)
	copyMap := newStringIntMap.Copy()
	fmt.Printf("%v -> %p\n", copyMap, copyMap)
	fmt.Println("<================================>")

	fmt.Println("проверка наличия ключа two:", newStringIntMap.Exists("two"))
	fmt.Println("проверка наличия ключа one:", newStringIntMap.Exists("one"))
	fmt.Println("<================================>")

	fmt.Println(newStringIntMap.Get("two"))
	fmt.Println(newStringIntMap.Get("one")) // значение будет дефолтное
}
