package main

import "fmt"

func main() {
	minScopes := 275

	var scopes1 int
	var scopes2 int
	var scopes3 int

	fmt.Println("Баллы ЕГЭ.")

	fmt.Println("Введите результат первого экзамена:")
	fmt.Scan(&scopes1)

	fmt.Println("Введите результат второго экзамена:")
	fmt.Scan(&scopes2)

	fmt.Println("Введите результат третьего экзамена:")
	fmt.Scan(&scopes3)

	fmt.Println("Сумма проходных баллов:", minScopes)

	totalScopes := scopes1 + scopes2 + scopes3
	fmt.Println("Количество набранных баллов:", totalScopes)

	if totalScopes >= minScopes {
		fmt.Println("Поздравляем! Вы поступили в институт.")
	} else {
		fmt.Println("Вы не поступили.")
	}

}
