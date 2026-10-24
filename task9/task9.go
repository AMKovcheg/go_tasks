package main

import "fmt"

func main() {
	var balance float64
	var choice int
	var amount float64

	fmt.Print("Введите начальный баланс: ")
	fmt.Scan(&balance)

	for {
		fmt.Println("\n=== Мини-банкомат ===")
		fmt.Println("1. Посмотреть баланс")
		fmt.Println("2. Пополнить баланс")
		fmt.Println("3. Снять средства")
		fmt.Println("0. Выйти")
		fmt.Print("Выберите действие: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Текущий баланс: %.2f\n", balance)
		case 2:
			fmt.Print("Введите сумму для пополнения: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Сумма должна быть положительной!")
			} else {
				balance += amount
				fmt.Printf("Баланс пополнен. Текущий баланс: %.2f\n", balance)
			}
		case 3:
			fmt.Print("Введите сумму для снятия: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Сумма должна быть положительной!")
			} else if amount > balance {
				fmt.Println("Недостаточно средств на балансе!")
			} else {
				balance -= amount
				fmt.Printf("Средства сняты. Текущий баланс: %.2f\n", balance)
			}
		case 0:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неверный выбор!")
		}
	}
}