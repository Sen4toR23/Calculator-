package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var romanNumerals = []RomanNumeral{
	{1, "I"},
	{4, "IV"},
	{5, "V"},
	{9, "IX"},
	{10, "X"},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	expression, _ := reader.ReadString('\n')
	result, err := evaluateExpression(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат:", result)
}

func evaluateExpression(expression string) (interface{}, error) {
	expression = strings.TrimSpace(expression)
	tokens := strings.Split(expression, " ")
	if len(tokens) != 3 {
		return nil, fmt.Errorf("некорректный формат выражения")
	}

	firstOperand, err := parseOperand(tokens[0])
	if err != nil {
		return nil, err
	}

	operator := tokens[1]

	secondOperand, err := parseOperand(tokens[2])
	if err != nil {
		return nil, err
	}

	switch operator {
	case "+":
		return firstOperand + secondOperand, nil
	case "-":
		return firstOperand - secondOperand, nil
	case "*":
		return firstOperand * secondOperand, nil
	case "/":
		if secondOperand == 0 {
			return nil, fmt.Errorf("деление на ноль")
		}
		return firstOperand / secondOperand, nil
	default:
		return nil, fmt.Errorf("недопустимая операция: %s", operator)
	}
}

func parseOperand(operand string) (int, error) {
	operand = strings.TrimSpace(operand)

	isRoman := false
	for _, romanNum := range romanNumerals {
		if operand == romanNum.Symbol {
			isRoman = true
			break
		}
	}

	if isRoman {
		value, err := parseRomanNumeral(operand)
		if err != nil {
			return 0, err
		}
		return value, nil
	} else {
		value, err := strconv.Atoi(operand)
		if err != nil {
			return 0, fmt.Errorf("некорректный операнд: %s", operand)
		}
		if value < 1 || value > 10 {
			return 0, fmt.Errorf("операнд должен быть числом от 1 до 10")
		}
		return value, nil
	}
}

func parseRomanNumeral(roman string) (int, error) {
	value := 0
	for len(roman) > 0 {
		symbolFound := false
		for _, numeral := range romanNumerals {
			if strings.HasPrefix(roman, numeral.Symbol) {
				value += numeral.Value
				roman = roman[len(numeral.Symbol):]
				symbolFound = true
				break
			}
		}
		if !symbolFound {
			return 0, fmt.Errorf("некорректное римское число: %s", roman)
		}
	}
	return value, nil
}
