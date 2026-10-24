package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введите число N: ")
	fmt.Scan(&n)

	sum := 0

	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
		}
	}

	fmt.Printf("Сумма чётных чисел от 1 до %d, делящихся на 3: %d\n", n, sum)
}