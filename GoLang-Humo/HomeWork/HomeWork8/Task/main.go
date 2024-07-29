package main

import (
	"fmt"
)

func main() {
	//var n int
	//fmt.Println("Вводите размер массива:")
	//fmt.Scan(&n)
	var arr []int
	var a int
	for i := 0; i < 5; i++ {
		fmt.Printf("Вводите значение a[%d] = ", i+1)
		fmt.Scan(&a)
		arr = append(arr, a)
	}
	//1
	fmt.Printf("Максимальный элемент в массиве: %d\n", max_arr(arr))
	//2
	fmt.Printf("Минимальный элемент в массиве: %d\n", min_arr(arr))
	//3
	fmt.Printf("Количество положительных чисел в массиве: %d\n", PodschPol(arr))
	//4
	fmt.Printf("Сумму всех элементов массива: %d\n", Sum_arr(arr))
	//5
	fmt.Printf("Среднее значение всех элементов массива: %.2f\n", SredZnach_arr(arr))
	//6
	var f int
	fmt.Println("Наш массив", arr)
	fmt.Println("Вводите число которую вы хотите удалить из массива")
	fmt.Scan(&f)
	fmt.Println("Удалили все вхожденый заданные числа из массива.", del_number(arr, f))
	//7
	var k int
	fmt.Println("На какое число умножить массив:")
	fmt.Scan(&k)
	new_arr := []int{}
	new_arr = umno(arr, k)
	fmt.Println("Новый массив: ", new_arr)
	//8
	var m int
	fmt.Println("Найти все индексы заданного числа в массиве.\nВводите число:")
	fmt.Scan(&m)
	fmt.Println("Все индексы заданного числа в массиве:\n", find_index(arr, m))
	//9
	fmt.Println("Копию массива: \n", copy_arr(arr))
	//10
	fmt.Println("Объядинение двух массивов: ", arr, new_arr)
	fmt.Println(will_unite(arr, new_arr))
	//11
	fmt.Println("Наш массив: ", arr)
	fmt.Println("Поменяли местами максимальный и минимальный элементы массива.", swap(arr))
	//12
	fmt.Println("Наш массив: ", arr)
	fmt.Printf("Является ли наш массив палиндромом? %s\n", palindrome(arr))
	//13
	fmt.Println("Второе наибольшее число в массиве равен: ", second_max(arr))
	//14
	fmt.Println("Наш массив: ", arr)
	fmt.Println("Перевернутый массив: ", reverse_arr(arr))
	//15
	fmt.Println("Наш массив: ", arr)
	fmt.Println("Удалили из массива все дубликаты", duplicate_arr(arr))
	//16
	fmt.Println("Наш массив: ", arr)
	fmt.Println("Переместили все нули в конец массива, сохроняя порядок ненулевых элементов ", MoveZero_arr(arr))
	//17
	var arr2 []int
	var b int
	for i := 0; i < 5; i++ {
		fmt.Printf("Вводите значение a[%d] = ", i+1)
		fmt.Scan(&b)
		arr2 = append(arr2, b)
	}
	fmt.Println("Первый массив: ", arr)
	fmt.Println("Второй массив: ", arr2)
	fmt.Println("Пересечение двух массивов: ", intersection_arr(arr, arr2))
	//18-----------------------
	//19
	fmt.Println("Первый массив: ", arr)
	fmt.Println("Второй массив: ", arr2)
	fmt.Println("Объединить двух отсортированных массивов в один отсортированный: ", SortTwo_arr(arr, arr2))
	//20
	var arr3 []int
	var c int
	for i := 0; i < 5; i++ {
		fmt.Printf("Вводите значение a[%d] = ", i+1)
		fmt.Scan(&c)
		arr3 = append(arr3, c)
	}
	fmt.Println("Наш массив: ", arr3)
	fmt.Println("Длина самого длиного подмасива, в котором все элементы различны равен: ", podArr(arr3))
}

// 1Найти максимальный элемент в массиве.
func max_arr(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// 2Найти минимальный элемент в массиве.
func min_arr(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

// 3Подсчитать количество положительных чисел в массиве.
func PodschPol(arr []int) int {
	var l int = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			l++
		}
	}
	return l
}

// 4Найти сумму всех элементов массива.
func Sum_arr(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

// 5Найти среднее значение всех элементов массива.
func SredZnach_arr(arr []int) float64 {
	s := float64(Sum_arr(arr))
	return s / float64(len(arr))
}

// 6Удалить все вхождения заданного числа из массива.
func del_number(arr []int, n int) []int {
	new_arr := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] != n {
			new_arr = append(new_arr, arr[i])
		}
	}
	return new_arr
}

// 7Умножить все элементы массива на заданное число.
func umno(arr []int, k int) []int {
	new_arr := []int{}
	for i := 0; i < len(arr); i++ {
		new_arr = append(new_arr, arr[i]*k)
	}
	return new_arr
}

// 8Найти все индексы заданного числа в массиве.
func find_index(arr []int, k int) []int {
	index := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] == k {
			index = append(index, i)
		}
	}
	return index
}

// 9Создать копию массива.
func copy_arr(arr []int) []int {
	new_arr := []int{}
	//copy(new_arr, arr)
	for i := 0; i < len(arr); i++ {
		new_arr = append(new_arr, arr[i])
	}
	return new_arr
}

// 10Объединить два массива.
func will_unite(arr []int, new_arr []int) []int {
	unite_arr := []int{}
	for i := 0; i < len(arr); i++ {
		unite_arr = append(unite_arr, arr[i])
	}
	for i := 0; i < len(new_arr); i++ {
		unite_arr = append(unite_arr, new_arr[i])
	}
	return unite_arr
}

// 11Поменять местами максимальный и минимальный элементы массива.
func swap(arr []int) []int {
	max := arr[0]
	min := arr[0]
	var imax int
	var imin int
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			imax = i
		}
		if arr[i] < min {
			min = arr[i]
			imin = i
		}
	}
	m := arr[imax]
	arr[imax] = arr[imin]
	arr[imin] = m
	return arr
}

// 12Проверить, является ли массив палиндромом.
func palindrome(arr []int) string {
	size := len(arr)
	var k int = 0
	for i := 0; i < size/2; i++ {
		if arr[i] == arr[size-i-1] {
			k++
		}
	}
	if size/2 == k {
		return "Да"
	} else {
		return "Нет"
	}
}

// 13Найти второе наибольшее число в массиве.
func second_max(arr []int) int {
	max1 := max_arr(arr)
	max2 := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] != max1 && arr[i] > max2 {
			max2 = arr[i]
		}
	}
	return max2
}

// 14Перевернуть массив.
func reverse_arr(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		b := arr[i]
		arr[i] = arr[len(arr)-i-1]
		arr[len(arr)-i-1] = b
	}
	return arr
}

// 15Удалить дубликаты из массива.
func duplicate_arr(arr []int) []int {
	new_arr := []int{}
	for i := 0; i < len(arr); i++ {
		var k int = 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				k++
			}
		}
		if k == 1 {
			new_arr = append(new_arr, arr[i])
		}
	}
	return new_arr
}

// 16Переместить все нули в конце массива, сохраняя порядок ненулевых элементов.
func MoveZero_arr(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := i; j < len(arr)-1; j++ {
				b := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = b
			}
		}
	}
	return arr
}

// 17Найти пересечение двух массивов.
func intersection_arr(arr1 []int, arr2 []int) []int {
	new_arr := []int{}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				new_arr = append(new_arr, arr1[i])
			}
		}
	}
	return new_arr
}

//18Проверить, является ли массив подмножеством другого массива.

// 19Объединить два отсортированных массива в один отсортированный.
func SortTwo_arr(arr1 []int, arr2 []int) []int {
	new_arr := []int{}
	for i := 0; i < len(arr1); i++ {
		new_arr = append(new_arr, arr1[i])
	}
	for i := 0; i < len(arr2); i++ {
		new_arr = append(new_arr, arr2[i])
	}
	sort(new_arr)
	return new_arr
}
func sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				b := arr[i]
				arr[i] = arr[j]
				arr[j] = b
			}
		}
	}
}

// 20Найти длину самого длинного подмассива, в котором все элементы различны.
func podArr(arr []int) int {
	max := 0
	var l int
	for i := 0; i < len(arr); i++ {
		l = 0
		for j := 0; j < len(arr); j++ {
			if arr[i] != arr[j] {
				l++
			} else {
				if l != 0 {
					break
				}
			}
		}
		if l > max {
			max = l
		}

	}
	return max + 1
}

//21Найти все подмассивы, сумма которых равна заданному числу.

//22Найти пару элементов в массиве, сумма которых равна заданному числу.
//23Найти наименьший положительный элемент, отсутствующий в массиве.
//24Найти максимальную сумму подмассива с условием, что подмассив не должен содержать более двух различных элементов.
//25Найти максимальную длину подмассива, сумма элементов которого равна заданному числу.
//26Найти максимальное произведение трех чисел в массиве.
//27Найти подмассив с максимальной суммой.
//28Переместить все отрицательные числа в начало массива, сохраняя порядок остальных чисел.
//29Найти подмассив с наибольшей длиной, сумма элементов которого равна нулю.
//30Найти наибольший общий делитель всех элементов массива.
