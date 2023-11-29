package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	operations := map[string]func(a, b int) int{
		"+": func(a, b int) int {
			return a + b
		},
		"-": func(a, b int) int {
			return a - b
		},
		"*": func(a, b int) int {
			return a * b
		},
		"/": func(a, b int) int {
			return int(a / b)
		},
	}

	var isRoman bool = false

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if !strings.ContainsAny(text, "+-*/") {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
		return
	}

	parts := strings.Split(text, " ")
	if len(parts) != 3 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}

	operator := parts[1]
	number1, err := strconv.Atoi(parts[0])
	if err != nil {
		if strings.ContainsAny(parts[0], "IVX") {
			number1 = convertFromRoman(parts[0])
			isRoman = true
		} else {
			fmt.Println("Вывод ошибки, так как введенное число может быть целым числом или состоять только из римских цифр")
			return
		}
	}
	number2, err := strconv.Atoi(parts[2])
	if (isRoman && err == nil) || (!isRoman && err != nil) {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		return
	} else {
		if strings.ContainsAny(parts[0], "IVX") {
			number2 = convertFromRoman(parts[2])
			isRoman = true
		} else {
			fmt.Println("Вывод ошибки, так как введенное число может быть целым числом или состоять только из римских цифр")
			return
		}
	}

	if number1 < 1 || number1 > 10 || number2 < 1 || number2 > 10 {
		fmt.Println("Вывод ошибки, так как на вход принимаются числа от 1 до 10 включительно, не более.")
		return
	}

	result := operations[operator](number1, number2)
	if result < 1 && isRoman {
		fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
		return
	}
	if isRoman {
		fmt.Println(convertToRoman(result))
	} else {
		fmt.Println(result)
	}
}

func convertFromRoman(strNumber string) int {
	var result int
	var lastChr string
	for i := len(strNumber) - 1; i >= 0; i-- {
		curChr := string(strNumber[i])
		if curChr == "I" {
			if lastChr == "X" || lastChr == "V" {
				result -= 1
			} else {
				result += 1
			}
			lastChr = "I"
		} else if curChr == "V" {
			result += 5
			lastChr = "V"
		} else if curChr == "X" {
			result += 10
			lastChr = "X"
		}
	}
	return result
}

func convertToRoman(number int) string {
	conversions := []struct {
		digit string
		value int
	}{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}

	var result string
	for _, conversion := range conversions {
		for number >= conversion.value {
			result += conversion.digit
			number -= conversion.value
		}
	}
	return result
}
