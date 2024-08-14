package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	////Task 1
	//n, err := countCharacters("test.txt")
	//if err != nil {
	//	fmt.Println("Ощибка при чтений файла:", err)
	//}
	//fmt.Printf("Количество символов в файле: %d\n", n)

	////Task 2
	//n, err := countLines("test.txt")
	//if err != nil {
	//	fmt.Println("Ощибка при открыти файла:", err)
	//}
	//fmt.Printf("Количество строк в файле: %d\n", n)

	////Task 3
	//file, err := os.Open("test.txt")
	//if err != nil {
	//	fmt.Println("Ощибка при открыти файла:", err)
	//}
	//defer file.Close()
	//text, err := io.ReadAll(file)
	//if err != nil {
	//	fmt.Println("Ощибка при чтении файла:", err)
	//}
	//fmt.Printf("Количество слов в строке: %d\n", countWords(string(text)))

	////Task 4
	//fmt.Println("Ввлдите слова которую хотите записать в файл\n")
	//var text string
	//fmt.Scan(&text)
	//err := writeStringToFile("test.txt", text)
	//if err != nil {
	//	fmt.Println("Ощибка при записи слов в файл:", err)
	//}

	////Task 5
	//text, err := readFirstLine("test.txt")
	//if err != nil {
	//	fmt.Println("Ощибка при открытии файла:", err)
	//}
	//fmt.Printf("Первая строка в текстовом файле: %s\n", text)

	//Task 6
}

// 1	Подсчет символов в файле
// Описание: Напишите программу, которая считывает файл и подсчитывает количество символов в нём.
// Методы или функции:
// func countCharacters(fileName string) (int, error)
func countCharacters(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ощибка при открыти файла:", err)
		return 0, err
	}
	defer file.Close()
	buf := make([]byte, 1024)
	text, err := file.Read(buf)
	return utf8.RuneCount(buf[:text]), err
}

// 2	Подсчет строк в файле
// Описание: Напишите программу, которая считает количество строк в текстовом файле.
// Методы или функции:
func countLines(fileName string) (int, error) {
	file, err := os.Open(fileName)
	defer file.Close()
	l := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l++
	}
	return l, err
}

// 3	Счетчик слов в строке
// Описание: Напишите функцию, которая считает количество слов в строке.
// Методы или функции:
// func countWords(s string) int
func countWords(s string) int {
	arr := strings.Split(s, " ")
	return len(arr)
}

// 4	Запись строки в файл
// Описание: Напишите программу, которая записывает заданную строку в файл.
// Методы или функции:
// func writeStringToFile(fileName, content string) error
func writeStringToFile(fileName string, data string) error {
	err := os.WriteFile(fileName, []byte(data), 0666)
	if err != nil {
		return err
	}
	return err
}

// 5	Чтение первой строки файла
// Описание: Напишите программу, которая читает первую строку из текстового файла.
// Методы или функции:
// func readFirstLine(fileName string) (string, error)
func readFirstLine(fileName string) (string, error) {
	file, err := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text(), err
}

//6	Копирование содержимого одного файла в другой
//Описание: Напишите программу, которая копирует содержимое одного файла в другой.
//Методы или функции:
//func copyFile(src, dst string) error

//Чтение строки с консоли и запись в файл
//Описание: Напишите программу, которая читает строку с консоли и записывает её в файл.
//Методы или функции:
//func readAndWriteToFile(fileName string) error
//Обратное чтение файла
//Описание: Напишите программу, которая читает файл с конца в начало и выводит его содержимое на экран.
//Методы или функции:
//func reverseReadFile(fileName string) (string, error)
//Объединение содержимого двух файлов
//Описание: Напишите программу, которая объединяет содержимое двух файлов и записывает результат в новый файл.
//Методы или функции:
//func concatenateFiles(file1, file2, outputFile string) error
//Проверка существования файла
//Описание: Напишите функцию, которая проверяет, существует ли файл с заданным именем.
//Методы или функции:
//func fileExists(fileName string) bool
//Чтение и подсчет уникальных слов в файле
//Описание: Напишите программу, которая читает файл и подсчитывает количество уникальных слов.
//Методы или функции:
//func countUniqueWords(fileName string) (int, error)
//Поиск и замена слова в файле
//Описание: Напишите программу, которая заменяет все вхождения определенного слова в файле на другое слово.
//Методы или функции:
//func replaceWordInFile(fileName, oldWord, newWord string) error
//Подсчет частоты слов в файле
//Описание: Напишите программу, которая подсчитывает частоту каждого слова в файле.
//Методы или функции:
//func wordFrequency(fileName string) (map[string]int, error)
//Переворачивание строк в файле
//Описание: Напишите программу, которая переворачивает строки в файле и записывает их в новый файл.
//Методы или функции:
//func reverseLines(fileName, outputFile string) error
//Удаление пустых строк из файла
//Описание: Напишите программу, которая удаляет все пустые строки из файла.
//Методы или функции:
//func removeEmptyLines(fileName, outputFile string) error
//Сравнение двух файлов на идентичность
//Описание: Напишите программу, которая сравнивает два файла и определяет, идентичны ли они.
//Методы или функции:
//func compareFiles(file1, file2 string) (bool, error)
//Чтение файла с конца
//Описание: Напишите программу, которая читает файл с конца и выводит его содержимое на экран.
//Методы или функции:
//func readFileReverse(fileName string) error
//Подсчет количества строк с определенным словом
//Описание: Напишите программу, которая считает количество строк в файле, содержащих определенное слово.
//Методы или функции:
//func countLinesWithWord(fileName, word string) (int, error)
//Генерация файла с повторяющимися строками
//Описание: Напишите программу, которая генерирует файл, содержащий заданную строку, повторенную указанное количество раз.
//Методы или функции:
//func generateRepeatedLinesFile(fileName, line string, count int) error
//Подсчет количества символов в каждой строке файла
//Описание: Напишите программу, которая подсчитывает количество символов в каждой строке файла и выводит их на экран.
//Методы или функции:
//func countCharactersPerLine(fileName string) error
//Поиск самого длинного слова в файле
//Описание: Напишите программу, которая находит самое длинное слово в файле и выводит его на экран.
//Методы или функции:
//func findLongestWord(fileName string) (string, error)
//Подсчет частоты встречаемости символов в файле
//Описание: Напишите программу, которая подсчитывает частоту встречаемости каждого символа в файле.
//Методы или функции:
//func charFrequency(fileName string) (map[rune]int, error)
//Реверсирование слов в каждой строке файла
//Описание: Напишите программу, которая реверсирует слова в каждой строке файла и записывает результат в новый файл.
//Методы или функции:
//func reverseWordsInFile(fileName, outputFile string) error
//Подсчет количества слов в каждой строке файла
//Описание: Напишите программу, которая подсчитывает количество слов в каждой строке файла и выводит результат на экран.
//Методы или функции:
//func countWordsPerLine(fileName string) error
//Найти и заменить слово в файле, с сохранением регистра
//Описание: Напишите программу, которая находит и заменяет все вхождения слова в файле, сохраняя регистр (с учетом заглавных и строчных букв).
//Методы или функции:
//func replaceWordWithCase(fileName, oldWord, newWord string) error
//Поиск самого короткого слова в файле
//Описание: Напишите программу, которая находит самое короткое слово в файле и выводит его на экран.
//Методы или функции:
//func findShortestWord(fileName string) (string, error)
//Чтение и запись файла по частям
//Описание: Напишите программу, которая читает файл по частям (например, по 1024 байта) и записывает его содержимое в другой файл.
//Методы или функции:
//func copyFileInChunks(srcFileName, dstFileName string) error
//Подсчет количества символов, слов и строк в файле
//Описание: Напишите программу, которая подсчитывает количество символов, слов и строк в файле и выводит результат на экран.
//Методы или функции:
//func countFileStats(fileName string) (int, int, int, error)
//Поиск и удаление строки в файле
//Описание: Напишите программу, которая находит и удаляет строку с определенным содержимым из файла.
//Методы или функции:
//func deleteLineFromFile(fileName, lineToDelete string) error
//Сортировка строк в файле в алфавитном порядке
//Описание: Напишите программу, которая сортирует строки в файле в алфавитном порядке и записывает результат в новый файл.
//Методы или функции:
//func sortFileLines(fileName, outputFile string) error
