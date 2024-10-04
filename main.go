package main

import (
	"fmt"
	"unicode"
)

var values []int
var operators []byte

func applyOperation(operation byte, b, a int) int {
	switch operation {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		if b == 0 {
			fmt.Println("Error: Division by zero")
			return 0
		}
		return a / b
	case '^':
		result := 1
		for b > 0 {
			result *= a
			b--
		}
		return result
	}
	return 0
}

func precedence(operation byte) int {
	switch operation {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '^':
		return 3
	}
	return 0
}

func countValue() {
	if len(values) < 2 || len(operators) == 0 {
		fmt.Println("Error: Invalid expression")
		return
	}
	secondValue := values[len(values)-1]
	values = values[:len(values)-1]
	firstValue := values[len(values)-1]
	values = values[:len(values)-1]
	operation := operators[len(operators)-1]
	operators = operators[:len(operators)-1]
	values = append(values, applyOperation(operation, secondValue, firstValue))
}

func evaluate(expression string) int {
	values = []int{}
	operators = []byte{}

	i := 0
	for i < len(expression) {
		char := expression[i]
		if char == ' ' {
			i++
			continue
		} else if char == '(' {
			operators = append(operators, char)
		} else if unicode.IsDigit(rune(char)) {
			value := 0
			for i < len(expression) && unicode.IsDigit(rune(expression[i])) {
				value = value*10 + int(expression[i]-'0')
				i++
			}
			values = append(values, value)
			i-- // Возвращаемся на один символ назад, чтобы обработать текущий оператор
		} else if char == ')' {
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				countValue()
			}
			if len(operators) > 0 && operators[len(operators)-1] == '(' {
				operators = operators[:len(operators)-1] // Удаляем '('
			}
		} else if char == '+' || char == '-' || char == '*' || char == '/' || char == '^' {
			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(char) {
				countValue()
			}
			operators = append(operators, char)
		}
		i++
	}

	// Завершаем оставшиеся операции
	for len(operators) > 0 {
		countValue()
	}

	if len(values) == 1 {
		return values[0]
	}

	return 0
}

func userInterface() {
	var choice int
	var expression string
	for {
		fmt.Print("1 - Enter\n2 - Exit\nChoice: ")
		fmt.Scan(&choice)

		if choice == 2 {
			break
		} else if choice == 1 {
			fmt.Print("Input expression: ")
			fmt.Scan(&expression)
			result := evaluate(expression)
			fmt.Println("Result:", result)
		} else {
			fmt.Println("Invalid choice. Please select 1 or 2.")
		}
	}
}

func main() {
	userInterface()
}
