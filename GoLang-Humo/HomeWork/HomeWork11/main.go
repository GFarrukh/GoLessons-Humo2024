package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "Hello Farrukh"
	//b := "12"
	fmt.Println(GeneratingAllPodstrok(a))
}

// 1. Конкатенация строк
// Напишите функцию, которая принимает две строки и возвращает их конкатенацию.
func ConcatenationString(a string, b string) string {
	return a + b
}

// 2 Длина строки
// Напишите функцию, которая принимает строку и возвращает её длину.
func SizeString(a string) int {
	return len(a)
}

// 3 Проверка подстроки
// Напишите функцию, которая проверяет, содержит ли строка заданную подстроку.
func CheckPodstroka(a string, b string) string {
	l := strings.Contains(a, b)
	if l == true {
		return "Строка " + a + " содержить подстроку " + b
	} else {
		return "Строка " + a + " не содержить подстроку " + b
	}
}

// 4 Поиск подстроки
// Напишите функцию, которая находит индекс первого вхождения подстроки в строке.
// Верните -1, если подстрока не найдена.
func SearchPodstroka(a string, b string) int {
	l := strings.Contains(a, b)
	if l == true {
		return strings.Index(a, b)
	} else {
		return -1
	}
}

// 5 Замена подстроки
// Напишите функцию, которая заменяет все вхождения подстроки в строке на другую подстроку.
func Change_Podstroka(a string, b string, c string) string {
	return strings.Replace(a, b, c, -1) //-1 Это что????
}

// 6 Удаление пробелов
// Напишите функцию, которая удаляет пробелы в начале и в конце строки.
func DeleteSpaces(a string) string {
	return strings.TrimSpace(a)
}

// 7 Преобразование регистра
// Напишите функцию, которая преобразует строку к верхнему и нижнему регистру.
func ReformingRegister(a string) (string, string) {
	return strings.ToUpper(a), strings.ToLower(a)
}

// 8 Повторение строки
// Напишите функцию, которая повторяет строку заданное количество раз.
func RepeatString(a string, n int) string {
	l := ""
	for i := 0; i < n; i++ {
		l = l + a
	}
	return l
}

// 9 Разбиение строки
// Напишите функцию, которая разбивает строку по заданному разделителю и возвращает срез строк.
// 10Сравнение строк
// Напишите функцию, которая сравнивает две строки без учета регистра.
// 11Поиск и замена первой подстроки
// Напишите функцию, которая заменяет первое вхождение подстроки в строке на 	другую подстроку.

// 12 Инверсия строки
// Напишите функцию, которая переворачивает строку.
func TurnOverString(a string) string {
	n := strings.Split(a, "")
	for i := 0; i < len(n)/2; i++ {
		b := n[i]
		n[i] = n[len(n)-i-1]
		n[len(n)-i-1] = b
	}
	return strings.Join(n, "")
}

// 13 Подсчет символов
// Напишите функцию, которая подсчитывает количество каждого символа в строке и возвращает карту с результатами.

// 14Удаление символов
// Напишите функцию, которая удаляет все вхождения заданного символа из строки.
// 15Подсчет слов
// Напишите функцию, которая подсчитывает количество слов в строке. Словом 	считается последовательность символов, разделенная пробелами.
// 16Проверка префикса и суффикса
// Напишите функции, которые проверяют, начинается ли строка с заданного префикса и 	заканчивается ли строка заданным суффиксом.
// 17Удаление дубликатов символов
// Напишите функцию, которая удаляет дубликаты символов в строке.
// 18Форматирование строки
// Напишите функцию, которая форматирует строку, заменяя специальные символы 	HTML на их эквиваленты.
// 19Псевдонимы строк
// Напишите функцию, которая создает псевдоним для строки, заменяя пробелы и 	специальные символы подчеркиваниями.
// 20Разбор строки
// Напишите функцию, которая разбирает строку, содержащую числа, и возвращает их 	сумму.
// 21Обратный порядок слов
// Напишите функцию, которая принимает строку и возвращает строку с обратным 	порядком слов.
// 22Подсчет уникальных слов
// Напишите функцию, которая подсчитывает количество уникальных слов в строке.
// 23Поиск палиндромов
// Напишите функцию, которая проверяет, является ли строка палиндромом.

// 24Перемешивание слов
// Напишите функцию, которая случайным образом перемешивает слова в строке.
func RandomWorldInString(a string) {

}

// 25 Сортировка слов по длине
// Напишите функцию, которая сортирует слова в строке по их длине.
func SortWordInString(a string) string {
	s := strings.Split(a, " ")
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if len(s[i]) < len(s[j]) {
				b := s[i]
				s[i] = s[j]
				s[j] = b
			}
		}
	}
	return strings.Join(s, " ")
}

// 26 Генерация хеш-значения
//Напишите функцию, которая вычисляет хеш-значение строки с использованием 	алгоритма SHA-256.

// 27 Генерация всех подстрок
// Напишите функцию, которая генерирует все подстроки заданной строки.
func GeneratingAllPodstrok(a string) string {
	s := strings.Split(a, "")
	a_new := ""
	for i := 0; i < len(s)+1; i++ {
		for j := i + 1; j < len(s)+1; j++ {
			k := s[i:j]
			m := ""
			for l := range k {
				m = m + k[l]
			}
			a_new = a_new + m + " "
		}
	}
	return strings.TrimSpace(a_new)
}

// 28 Поиск всех анаграмм
// Напишите функцию, которая ищет все анаграммы строки в другой строке.
func SearchAnagram(a string, b string) {
}

// 29 Подсчет слов и символов
// Напишите функцию, которая подсчитывает количество слов и символов в строке.
func PodschetSlovSimvolov(a string) {
	s := strings.Split(a, " ")
	PodschetSlov := len(s)
	s1 := strings.Replace(a, " ", "", -1)
	PodschetSimvolov := len(s1)
	fmt.Println("Количество слов в", a, ":", PodschetSlov, "\nКоличество символов в", a, ":", PodschetSimvolov)
}
