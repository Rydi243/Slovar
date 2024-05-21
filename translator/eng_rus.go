package translator

import (
	"fmt"
)

func Eng_rus(numf int, eng_rus_maps map[string]string) {
	var word string
	if numf == 1 {
		fmt.Println("Введите слово на английском: ")
		fmt.Scan(&word)
		if eng_rus_maps[word] == "" {
			fmt.Println("Не знаю такого слова")
		} else {
			fmt.Print(word, " -> ", eng_rus_maps[word])
		}
	} else if numf == 2 {
		var wordeng string
		fmt.Print("Введите слово которое хотите добавить на английском: ")
		fmt.Scan(&word)
		fmt.Print("Введите перевод на русском: ")
		fmt.Scan(&wordeng)
		eng_rus_maps[word] = wordeng
		fmt.Println(word, " -> ", eng_rus_maps[word])
	} else if numf == 3 {
		fmt.Print("Введите слово которое хотите удалить: ")
		fmt.Scan(&word)
		if eng_rus_maps[word] == "" {
			fmt.Println("Такого слова и так нет")
		} else {
			delete(eng_rus_maps, word)
			fmt.Println("Слово удалено", " -> ", word)
		}
	} else if numf == 4 {
		for key, value := range eng_rus_maps {
			fmt.Println(key, " -> ", value)
		}
	}
}
