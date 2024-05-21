package main

import (
	"fmt"
	"task_go/translator"
)

func main() {
	var num, numf int
	rus_eng_maps := make(map[string]string)
	eng_rus_maps := make(map[string]string)
	num = 5

	for num != 0 {
		fmt.Println("1. Русско - Английский словарь")
		fmt.Println("2. Английско - Русский словарь")
		fmt.Println("0. Выход")
		fmt.Scan(&num)
		numf = 5

		for numf != 0 {
			if num == 1 {
				fmt.Println("1. Поиск")
				fmt.Println("2. Добавить")
				fmt.Println("3. Удалить")
				fmt.Println("4. Показать все")
				fmt.Println("0. Назад")
				fmt.Scan(&numf)

				if numf == 1 || numf == 2 || numf == 3 || numf == 4 {
					translator.Rus_eng(numf, rus_eng_maps)
				} else if numf == 0 {
					break
				}

			} else if num == 2 {
				fmt.Println("1. Поиск")
				fmt.Println("2. Добавить")
				fmt.Println("3. Удалить")
				fmt.Println("4. Показать все")
				fmt.Println("0. Назад")
				fmt.Scan(&numf)

				if numf == 1 || numf == 2 || numf == 3 || numf == 4 {
					translator.Eng_rus(numf, eng_rus_maps)
				} else if numf == 0 {
					break
				}
			} else if num == 0 {
				fmt.Println("Пока")
				break
			}
		}
	}
}
