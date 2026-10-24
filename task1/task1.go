package main

import "fmt"

func main() {
	var a, b float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Print("Введите операцию (+, -, *, /): ")
	fmt.Scan(&op)

	var result float64
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль!")
			return
		}
		result = a / b
	default:
		fmt.Println("Неизвестная операция!")
		return
	}

	fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", a, op, b, result)
}