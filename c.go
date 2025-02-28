package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите выражение:")
	if !scanner.Scan() {
		panic("Ошибка чтения ввода")
	}
	line, sym := strChoper(scanner.Text())

	if len(line) != 2 {
		panic("Должно быть два операнда")
	}
	if sym == "0" {
		panic("используйте один из следующих операторов: +, -, *, /")
	}

	validOperand1, operand1 := firstStrValid(line)
	validOperand2, operand2 := secondStrValid(line)

	if validOperand1 {
		panic("первый операнд должен быть строкой меньше 10 символов")
	}

	if !validOperand2 {
		panic("второй операнд должен быть числом меншье 10, либо строкой меньше 10 символов")
	}

	result := stringCalc(operand1, operand2, sym)
	if len(result) > 40 {
		result = result[:40] + "..."
		fmt.Printf("\"%s\"", result)
	} else {
		fmt.Printf("\"%s\"", result)
	}

}
func stringCalc(operand1, operand2, sym string) string {
	var result string // Объявляем result заранее

	switch sym {
	case "+":
		result = operand1 + operand2

	case "-":
		if strings.Contains(operand1, operand2) {
			result = strings.ReplaceAll(operand1, operand2, "")
		} else {
			result = operand1
		}

	case "*":
		num, err := strconv.Atoi(operand2)
		if err != nil {
			panic("Ошибка: при умножении строки второй операнд должен быть числом")
		}
		result = strings.Repeat(operand1, num)

	case "/":
		num, err := strconv.Atoi(operand2)
		if err != nil {
			panic("Ошибка: при делении строки второй операнд должен быть числом")
		}
		if num == 0 {
			panic("Ошибка: деление на ноль недопустимо")
		}
		if len(operand1)%num != 0 {
			panic("Ошибка: строка не делится на целое число")
		}
		result = operand1[:len(operand1)/num]

	default:
		panic("используйте один из следующих операторов: +, -, *, /")
	}

	return result

}

func strChoper(line string) ([]string, string) {
	var symbol string
	var words []string

	if strings.Contains(line, "+") {
		words = strings.Split(line, " + ")
		symbol = "+"
	} else if strings.Contains(line, "-") {
		words = strings.Split(line, " - ")
		symbol = "-"
	} else if strings.Contains(line, "*") {
		words = strings.Split(line, " * ")
		symbol = "*"
	} else if strings.Contains(line, "/") {
		words = strings.Split(line, " / ")
		symbol = "/"
	} else {
		words = nil
		symbol = "0"
	}

	return words, symbol
}

func firstStrValid(line []string) (bool, string) {
	operand1 := line[0]

	if len(line) == 1 {
		return false, operand1[1 : len(operand1)-1]
	}

	_, err := strconv.Atoi(operand1)
	if err == nil {
		return true, operand1[1 : len(operand1)-1] // Если operand1 - число, возвращаем true
	}

	if len(operand1) > 12 { // Если строка длиннее 12 символов (отсекли 2), то true
		return true, operand1[1 : len(operand1)-1]
	}

	return false, operand1[1 : len(operand1)-1]
}

func secondStrValid(line []string) (bool, string) {
	operand2 := line[1]

	// Проверяем, является ли operand2 числом
	num2, err := strconv.Atoi(operand2)
	if err == nil && num2 <= 10 { // Если число и <= 10
		return true, operand2
	}

	// Проверяем, заключена ли строка в кавычки
	if strings.HasPrefix(operand2, "\"") && strings.HasSuffix(operand2, "\"") {
		stripped := operand2[1 : len(operand2)-1] // Убираем кавычки
		if len(stripped) < 10 {
			return true, stripped
		}
	}

	return false, operand2
}
