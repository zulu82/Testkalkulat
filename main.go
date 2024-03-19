package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Функция для конвертации римского числа в арабское
func romanToArabic(roman string) (int, error) {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}
	arabic := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[rune(roman[i])]
		if value < prevValue {
			arabic -= value
		} else {
			arabic += value
		}
		prevValue = value
	}
	return arabic, nil
}

// Функция для конвертации арабского числа в римское
func arabicToRoman(arabic int) (string, error) {
	if arabic < 1 {
		return "", fmt.Errorf("результат меньше единицы")
	}

	numerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}
	var result strings.Builder

	for _, numeral := range numerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String(), nil
}

// Основная функция калькулятора
func calculate(expression string) (string, error) {
	// Проверка на соответствие формату
	match, _ := regexp.MatchString(`^(?:[1-9]|10|[IVXLCDM]+) [+\-*/] (?:[1-9]|10|[IVXLCDM]+)$`, expression)
	if !match {
		return "", fmt.Errorf("неверный формат ввода")
	}

	// Разделение ввода на компоненты
	parts := strings.Split(expression, " ")
	a, b := parts[0], parts[2]
	operator := parts[1]

	// Определение типа чисел и конвертация при необходимости
	isRoman := false
	if _, err := strconv.Atoi(a); err != nil {
		isRoman = true
		aInt, err := romanToArabic(a)
		if err != nil {
			return "", err
		}
		a = strconv.Itoa(aInt)

		bInt, err := romanToArabic(b)
		if err != nil {
			return "", err
		}
		b = strconv.Itoa(bInt)
	}

	// Выполнение операции
	aInt, _ := strconv.Atoi(a)
	bInt, _ := strconv.Atoi(b)
	var result int
	switch operator {
	case "+":
		result = aInt + bInt
	case "-":
		result = aInt - bInt
	case "*":
		result = aInt * bInt
	case "/":
		if bInt == 0 {
			return "", fmt.Errorf("деление на ноль")
		}
		result = aInt / bInt
	default:
		return "", fmt.Errorf("неверный оператор")
	}

	// Возврат результата в нужном формате
	if isRoman {
		return arabicToRoman(result)
	}
	return strconv.Itoa(result), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	result, err := calculate(expression)
	if err != nil {
		panic(err)
	}
	fmt.Println("Результат:", result)
}
