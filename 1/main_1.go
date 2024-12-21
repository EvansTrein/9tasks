package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

// 1. Создает несколько переменных различных типов данных
var numDecimal int = 42           // Десятичная система
var numOctal int = 052            // Восьмеричная система
var numHexadecimal int = 0x2A     // Шестнадцатиричная система
var pi float64 = 3.14             // Тип float64
var name string = "Golang"        // Тип string
var isActive bool = true          // Тип bool
var complexNum complex128 = 1 + 2i // Тип complex128

func DefinitionType(variable interface{}) string {
    t := fmt.Sprintf("Тип переменной: %T", variable)
	return t
}

func main() {

	// 2. Определяет тип каждой переменной и выводит его на экран
	fmt.Printf("numDecimal -> %v, type %T\n", numDecimal, numDecimal)
	fmt.Printf("numOctal -> %v, type %T\n", numOctal, numOctal)
	fmt.Printf("numHexadecimal -> %v, type %T\n", numHexadecimal, numHexadecimal)
	fmt.Printf("pi -> %v, type %T\n", pi, pi)
	fmt.Printf("name -> %v, type %T\n", name, name)
	fmt.Printf("isActive -> %v, type %T\n", isActive, isActive)
	fmt.Printf("complexNum -> %v, type %T\n", complexNum, complexNum)
	fmt.Println("<================================>")

	// 3. Преобразует все переменные в строковый тип и объединяет их в одну строку
	numDecimalStr := strconv.Itoa(numDecimal)
	numOctalStr := fmt.Sprintf("%03o", numOctal)            //  strconv.FormatInt(int64(numOctal), 8) тут 0 не будет в строке
	numHexadecimalStr := fmt.Sprintf("%#x", numHexadecimal) // strconv.FormatInt(int64(numHexadecimal), 16)
	piStr := strconv.FormatFloat(pi, 'f', 2, 64)
	isActiveStr := strconv.FormatBool(isActive)
	complexNumStr := fmt.Sprintf("%g+%gi", real(complexNum), imag(complexNum))
	oneStr := numDecimalStr + numOctalStr + numHexadecimalStr + piStr + name + isActiveStr + complexNumStr
	fmt.Println(oneStr)
	fmt.Println("<================================>")

	// 4. Преобразовать эту строку в срез рун
	runes := []rune(oneStr)
	fmt.Println(runes)

	// 5. Захэшировать этот срез рун SHA256, добавив в середину соль "go-2024" и вывести результат
	salt := []byte("go-2024")
	runesBytesLeft := []byte(string(runes[:len(runes)/2]))
	runesBytesRight := []byte(string(runes[len(runes)/2:]))
	allBytesSlice := make([]byte, 0, len(runesBytesLeft)+len(runesBytesRight)+len(salt))
	allBytesSlice = append(allBytesSlice, runesBytesLeft...)
	allBytesSlice = append(allBytesSlice, salt...)
	allBytesSlice = append(allBytesSlice, runesBytesRight...)

	hash := sha256.Sum256(allBytesSlice)
	fmt.Println(hex.EncodeToString(hash[:]))

}
