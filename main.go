package main

import (
	"fmt"
	"task_go/translator"
)

func printMenu() {
	fmt.Println("1. Русско - Английский словарь")
	fmt.Println("2. Английско - Русский словарь")
	fmt.Println("0. Выход")
}

func printSubMenu() {
	fmt.Println("1. Поиск")
	fmt.Println("2. Добавить")
	fmt.Println("3. Удалить")
	fmt.Println("4. Показать все")
	fmt.Println("0. Назад")
}

func menu_one(m map[string]string) {
	for {
		num := 0
		printSubMenu()
		fmt.Scan(&num)

		if num == 1 || num == 2 || num == 3 || num == 4 {
			translator.Rus_eng(num, m)
		} else if num == 0 {
			break
		}
	}
}

func menu_two(m map[string]string) {
	for {
		num := 0
		printSubMenu()
		fmt.Scan(&num)

		if num == 1 || num == 2 || num == 3 || num == 4 {
			translator.Eng_rus(num, m)
		} else if num == 0 {
			break
		}
	}
}

func main() {
	rus_eng_maps := make(map[string]string)
	eng_rus_maps := make(map[string]string)

	for {
		num := 0
		printMenu()
		fmt.Scan(&num)

		switch num {
		case 0:
			fmt.Println("Пока")
			return
		case 1:
			menu_one(rus_eng_maps)
		case 2:
			menu_two(eng_rus_maps)
		}
	}
}
