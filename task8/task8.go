package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	if isPrime(n) {
		fmt.Printf("Число %d является простым\n", n)
	} else {
		fmt.Printf("Число %d не является простым\n", n)
	}
}