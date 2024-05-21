package translator

import (
	"fmt"
)

func Rus_eng(numf int, rus_eng_maps map[string]string) {
	var word string
	if numf == 1 {
		fmt.Print("Введите слово на русском: ")
		fmt.Scan(&word)
		if rus_eng_maps[word] == "" {
			fmt.Println("Не знаю такого слова")
		} else {
			fmt.Println(word, " -> ", rus_eng_maps[word])
		}
	} else if numf == 2 {
		var wordeng string
		fmt.Print("Введите слово которое хотите добавить на русском: ")
		fmt.Scan(&word)
		fmt.Print("Введите перевод на английском: ")
		fmt.Scan(&wordeng)
		rus_eng_maps[word] = wordeng
		fmt.Println(word, " -> ", rus_eng_maps[word])
	} else if numf == 3 {
		fmt.Print("Введите слово которое хотите удалить: ")
		fmt.Scan(&word)
		if rus_eng_maps[word] == "" {
			fmt.Println("Такого слова и так нет")
		} else {
			delete(rus_eng_maps, word)
			fmt.Println("Слово удалено", " -> ", word)
		}
	} else if numf == 4 {
		for key, value := range rus_eng_maps {
			fmt.Println(key, " -> ", value)
		}
	}
}
