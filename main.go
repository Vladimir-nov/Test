package main

import "fmt"

func input_num() uint {
	fmt.Print("Введите число для вычисления его факториала: ")
	var num float64
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Ошибка ввода.")
		return input_num()
	} else if num < 0 {
		fmt.Println("Число должно быть положительным.")
		return input_num()
	}
	return uint(num)
}

func factorial(num uint) uint {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
func main() {
	fmt.Print("Факториал числа равен: ", factorial(input_num()))
}
