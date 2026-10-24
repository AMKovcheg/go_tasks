package main

import "fmt"

func main() {
	var n int

	fmt.Print("Введите количество чисел N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("N должно быть больше 0")
		return
	}

	sum := 0
	evenCount := 0
	oddCount := 0
	var max, min int

	for i := 1; i <= n; i++ {
		var num int
		fmt.Printf("Введите число %d: ", i)
		fmt.Scan(&num)

		sum += num

		if i == 1 {
			max = num
			min = num
		} else {
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}

		if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	avg := float64(sum) / float64(n)

	fmt.Println("\n=== Статистика ===")
	fmt.Printf("Сумма всех чисел: %d\n", sum)
	fmt.Printf("Среднее значение: %.2f\n", avg)
	fmt.Printf("Максимальное значение: %d\n", max)
	fmt.Printf("Минимальное значение: %d\n", min)
	fmt.Printf("Количество чётных чисел: %d\n", evenCount)
	fmt.Printf("Количество нечётных чисел: %d\n", oddCount)
}