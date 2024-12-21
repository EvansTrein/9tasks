package main

import "testing"

func TestStringIntMap(t *testing.T) {
	testStruct := NewStringIntMap()

	t.Run("AddTest", func(t *testing.T) {
		testStruct.Add("one", 1)
		testStruct.Add("two", 2)
		testStruct.Add("three", 3)

		if testStruct.data["one"] != 1 {
			t.Errorf("ключ one с значением 1 не был добавлен")
		}

		if testStruct.data["two"] != 2 {
			t.Errorf("ключ two с значением 2 не был добавлен")
		}

		if testStruct.data["three"] != 3 {
			t.Errorf("ключ three с значением 3 не был добавлен")
		}
	})

	t.Run("RemoveTest", func(t *testing.T) {
		testStruct.Remove("three")
		if _, ok := testStruct.data["three"]; ok {
			t.Errorf("ключ three не был удален")
		}
	})

	t.Run("CopyTest", func(t *testing.T) {
		copied := testStruct.Copy()

		if len(copied) != len(testStruct.data) {
			t.Errorf("количество элементов после копирования не совпадает")
		}

		for key := range copied {
			if copied[key] != testStruct.data[key] {
				t.Errorf("после копирования значения ключей не совпадают")
			}
		}
	})

	t.Run("ExistsTest", func(t *testing.T) {

		if exists := testStruct.Exists("two"); !exists {
			t.Errorf("ключа two нет в мапе, хотя должен быть")
		}

		if exists := testStruct.Exists("three"); exists {
			t.Errorf("ключ three есть в мапе, но его там быть не должно")
		}
	})

	t.Run("GetTest", func(t *testing.T) {
		value, ok :=  testStruct.Get("one")

		if !ok {
			t.Errorf("ключа one нет в мапе, хотя должен быть")
		}
		if value != testStruct.data["one"] {
			t.Errorf("полученное значение по ключу one не совпадает с значением по этому же ключу в структуре")
		}
	})
}
