package main

import (
	"fmt"
	"log"
)

func main() {
	const minScopes = 275
	const limit = 100

	var scopes1 int
	var scopes2 int
	var scopes3 int

	fmt.Println("Баллы ЕГЭ.")

	fmt.Println("Введите результат первого экзамена:")
	fmt.Scan(&scopes1)
	if scopes1 < 0 {
		fmt.Println("Значение баллов не может быть отрицательным")
		log.Fatal()
	}
	if scopes1 > limit {
		fmt.Println("Значение баллов не может превышать", limit, "баллов")
		log.Fatal()
	}

	fmt.Println("Введите результат второго экзамена:")
	fmt.Scan(&scopes2)
	if scopes2 < 0 {
		fmt.Println("Значение баллов не может быть отрицательным")
		log.Fatal()
	}
	if scopes2 > limit {
		fmt.Println("Значение баллов не может превышать", limit, "баллов")
		log.Fatal()
	}

	fmt.Println("Введите результат третьего экзамена:")
	fmt.Scan(&scopes3)
	if scopes3 < 0 {
		fmt.Println("Значение баллов не может быть отрицательным")
		log.Fatal()
	}
	if scopes3 > limit {
		fmt.Println("Значение баллов не может превышать", limit, "баллов")
		log.Fatal()
	}

	fmt.Println("Сумма проходных баллов:", minScopes)

	totalScopes := scopes1 + scopes2 + scopes3
	fmt.Println("Количество набранных баллов:", totalScopes)

	if totalScopes >= minScopes {
		fmt.Println("Поздравляем! Вы поступили в институт.")
	} else {
		fmt.Println("Вы не поступили.")
	}
}
