package main

import "fmt"

//Поворот  массива в лево
func turnLeft() {
	var (
		countArr, count int
	)
	fmt.Scanln(&countArr)
	// создаем slice
	arr := make([]int, countArr)
	for i := 0; i < countArr; i++ {
		fmt.Scanln(&arr[i])
	}
	fmt.Scanln(&count) //количество поворотов
	// поворот
	for j := 1; j <= count%countArr; j++ {
		a := arr[len(arr)-1]
		for i := len(arr) - 1; i > 0; i-- {
			arr[i] = arr[i-1]
		}
		arr[0] = a
	}
	fmt.Println(arr)
}

//Поворот  массива в право
func turnRight() {
	var (
		arrS            = [1]int{1}
		countArr, count int
	)
	fmt.Scanln(&countArr) //Количество элементов массива
	// обнуляем и пересоздаем массив
	arr := arrS[:0]
	for i, b := 0, 0; i <= countArr-1; i++ {
		fmt.Scanln(&b)
		arr = append(arr, b)
	}
	fmt.Scanln(&count) //количество поворотов
	// поворот
	for j := 1; j <= count; j++ {
		a := arr[0]
		for i := 0; i < len(arr)-1; i++ {
			arr[i] = arr[i+1]
		}
		arr[len(arr)-1] = a
	}
	fmt.Println(arr)
}
