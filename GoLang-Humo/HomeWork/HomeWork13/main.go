package main

import (
	"fmt"
	"strings"
)

func main() {
	//text1 := MyString{str: "FGDDc dfvdf dfvdf"}
	//fmt.Println(text1)
	//PrintTask1(text1)
	//text2 := MyFormatter{str: "Goooooooo"}
	//fmt.Println(text2)
	//PrintTask2(text2)
	//text3 := IntPointer{num: 10}
	//fmt.Println(text3)
	//PrintTask3(text3)
	//text4 := TextCleaner{text: " My name is Farrukh "}
	//fmt.Println(text4)
	//PrintTask4(text4)
	//var a []string
	//a = append(a, "a")
	//a = append(a, "b")
	//a = append(a, "c")
	//a = append(a, "d")
	//a = append(a, "e")
	//a = append(a, "f")
	//a = append(a, "g")
	//a = append(a, "h")
	//a = append(a, "i")
	//a = append(a, "j")
	//a = append(a, "k")
	//a = append(a, "l")
	//text5 := StringJoiner{arr: a}
	//separator := "125"
	//fmt.Println(text5)
	//PrintTask5(text5, separator)
	//text6 := TextExtractor{arr: a}
	//start := 3
	//length1 := 5
	//length2 := 4
	//fmt.Println(text6)
	//PrintTask6(text6, start, length1, length2)
	text7 := TextReverser{str: "My name is Farrukh"}
	fmt.Println(text7)
	PrintTask7(text7)
}

// 1.  Длина строки и Количество слов
// Описание: Реализуйте интерфейс StringProcessor, который будет содержать методы Length() и WordCount().
// Реализуйте структуру MyString, которая будет работать с текстом и реализует этот интерфейс.
// Методы:
// Length() int для получения длины строки.
// WordCount() int для подсчета количества слов
type StringProcessor interface {
	Length() int
	WordCount() int
}
type MyString struct {
	str string
}

func (n MyString) Length() int {
	return len(n.str)
}
func (n MyString) WordCount() int {
	l := strings.Split(n.str, " ")
	return len(l)
}
func PrintTask1(s StringProcessor) {
	fmt.Printf("Длина строк: %d, Количество слов в строке: %d\n", s.Length(), s.WordCount())
}

// 2. Форматирование строки
// Описание: Реализуйте интерфейс Formatter с методами ToUpper() и ToLower().
// Реализуйте структуру MyFormatter, которая будет работать со строкой и реализует этот интерфейс.
// Методы:
// ToUpper() string для преобразования строки в верхний регистр.
// ToLower() string для преобразования строки в нижний регистр.
type Formatter interface {
	ToUpper() string
	ToLower() string
}
type MyFormatter struct {
	str string
}

func (n MyFormatter) ToUpper() string {
	return strings.ToUpper(n.str)
}
func (n MyFormatter) ToLower() string {
	return strings.ToLower(n.str)
}
func PrintTask2(s Formatter) {
	fmt.Printf("Преоброзование строки в верхний регистор: %s\nПреоброзование строки в нижний регистр: %s\n", s.ToUpper(), s.ToLower())
}

// 3. Работа с указателями на числа
// Описание: Реализуйте интерфейс PointerOperations с методами Increment() и Decrement().
// Реализуйте структуру IntPointer с указателем на число, которая реализует этот интерфейс.
// Методы:
// Increment() для увеличения значения числа на 1.
// Decrement() для уменьшения значения числа на 1.
type PointerOperations interface {
	Increment() int64
	Decrement() int64
}
type IntPointer struct {
	num int64
}

func (n IntPointer) Increment() int64 {
	n.num++
	return n.num
}
func (n IntPointer) Decrement() int64 {
	n.num--
	return n.num
}
func PrintTask3(s PointerOperations) {
	fmt.Printf("Число увеличеное на 1: %d\nЧисло уменьшеное на 1: %d\n", s.Increment(), s.Decrement())
}

// 4. Удаление пробелов в строке
// Описание: Реализуйте интерфейс StringCleaner с методами TrimSpaces() и RemoveSpaces().
// Реализуйте структуру TextCleaner, которая будет работать со строками и реализует этот интерфейс.
// Методы:
// TrimSpaces() string для удаления пробелов с начала и конца строки.
// RemoveSpaces() string для удаления всех пробелов из строки.
type StringCleaner interface {
	TrimSpaces() string
	RemoveSpaces() string
}
type TextCleaner struct {
	text string
}

func (n TextCleaner) TrimSpaces() string {
	return strings.TrimSpace(n.text)
}
func (n TextCleaner) RemoveSpaces() string {
	return strings.Replace(n.text, " ", "", -1)
}
func PrintTask4(s StringCleaner) {
	fmt.Printf("Наша новая строка после удаления пробелов с начало и конца старой строки: %s\n", s.TrimSpaces())
	fmt.Printf("Наша новая строка после удоления всех пробелов из строки: %s\n", s.RemoveSpaces())
}

// 5. Конкатенация строк
// Описание: Реализуйте интерфейс StringConcatenator с методами Concat() и Join().
// Реализуйте структуру StringJoiner, которая будет работать с массивами строк и реализует этот интерфейс.
// Методы:
// Concat() string для конкатенации всех строк в массиве.
// Join(separator string) string для объединения всех строк с заданным разделителем.
type StringConcatenator interface {
	Concat() string
	Join() string
}
type StringJoiner struct {
	arr []string
}

func (s StringJoiner) Concat() string {
	l := ""
	for i := range s.arr {
		l = l + s.arr[i]
	}
	return l
}
func (s StringJoiner) Join(separator string) []string {
	for i := range s.arr {
		s.arr[i] = s.arr[i] + separator
	}
	return s.arr
}
func PrintTask5(s StringJoiner, separator string) {
	fmt.Printf("Конкатенция всех строк в массиве: %s\n", s.Concat())
	fmt.Printf("Объедение всех строк с заданным разделителем %s : %s\n", separator, s.Join(separator))
}

// 6. Извлечение подстроки
// Описание: Реализуйте интерфейс SubstringExtractor с методами Substr(start, length int) string и Suffix(length int) string.
// Реализуйте структуру TextExtractor, которая работает со строками и реализует этот интерфейс.
// Методы:
// Substr(start, length int) string для извлечения подстроки заданной длины с указанной позиции.
// Suffix(length int) string для извлечения последних n символов строки.
type SubstringExtractor interface {
	Substr() string
	Suffix() string
}
type TextExtractor struct {
	arr []string
}

func (s TextExtractor) Substr(start, length int) []string {
	return s.arr[start:(start + length)]
}
func (s TextExtractor) Suffix(length int) []string {
	return s.arr[(len(s.arr) - length):]
}
func PrintTask6(s TextExtractor, start int, length1 int, length2 int) {
	fmt.Printf("Извлекли подстроку длиной %d с позиции %d : %s\n", length1, start, s.Substr(start, length1))
	fmt.Printf("Извлекли последные %d символов строки: %s", length2, s.Suffix(length2))
}

// 7. Переворот строки
// Описание: Реализуйте интерфейс StringReverser с методами Reverse() string и ReverseWords() string.
// Реализуйте структуру TextReverser, которая будет работать со строками и реализует этот интерфейс.
// Методы:
// Reverse() string для переворота строки.
// ReverseWords() string для переворота слов в строке.
type StringReverser interface {
	Reverse() string
	ReverseWords() string
}
type TextReverser struct {
	str string
}

func (text TextReverser) Reverse() string {
	arr := strings.Split(text.str, "")
	new_text := ""
	for i := len(arr) - 1; i >= 0; i-- {
		new_text += arr[i]
	}
	return new_text
}
func (text TextReverser) ReverseWords() string {
	arr := strings.Split(text.str, " ")
	new_text := ""
	for i := range arr {
		n := strings.Split(arr[i], "")
		for j := len(n) - 1; j >= 0; j-- {
			new_text = new_text + n[j]
		}

		new_text = new_text + " "
	}
	return new_text
}
func PrintTask7(s StringReverser) {
	fmt.Printf("Переворот строки: %s\n", s.Reverse())
	fmt.Printf("Переворот слов в строке: %s\n", s.ReverseWords())
}

//8. Вставка подстроки
//Описание: Реализуйте интерфейс StringInserter с методами InsertAt(index int, substring string) string
//и InsertAfterWord(word, substring string) string.
//Реализуйте структуру TextInserter, которая работает со строками и реализует этот интерфейс.
//Методы:
//InsertAt(index int, substring string) string для вставки подстроки в указанную позицию.
//InsertAfterWord(word, substring string) string для вставки подстроки после указанного слова.
type StringInserter interface {
	InsertAt() string
	InsertAfterWord() string
}
type TextInserter struct {
	text string
	index int
	substring string
	word string
}
func (s TextInserter) InsertAt(n TextInserter) string {
	arr := strings.Split(s.text, "")
	new_text := ""
	for i := range arr {
		new_text += arr[i]
		if i == n.index {
			new_text += n.substring
		}
	}
	return new_text
}
func (s TextInserter) InsertAfterWord(n TextInserter) string {

}

//9. Проверка на палиндром
//Описание: Реализуйте интерфейс PalindromeChecker с методами IsPalindrome() bool и IsWordPalindrome() bool. Реализуйте структуру TextPalindromeChecker, которая будет работать со строками и реализует этот интерфейс.
//Методы:
//IsPalindrome() bool для проверки, является ли строка палиндромом.
//IsWordPalindrome() bool для проверки, является ли строка палиндромом по словам.
//10. Удаление повторяющихся слов
//Описание: Реализуйте интерфейс DuplicateRemover с методами RemoveDuplicates() и RemoveDuplicatesCaseInsensitive(). Реализуйте структуру TextDuplicateRemover, которая работает со строками и реализует этот интерфейс.
//Методы:
//RemoveDuplicates() string для удаления дубликатов слов в строке.
//RemoveDuplicatesCaseInsensitive() string для удаления дубликатов слов без учета регистра.
//11. Бюджетирование
//Описание: Реализуйте интерфейс Budget с методами AddIncome(amount float64) и AddExpense(amount float64). Реализуйте структуру MonthlyBudget, которая будет работать с месячным бюджетом и реализует этот интерфейс.
//Методы:
//AddIncome(amount float64) для добавления дохода.
//AddExpense(amount float64) для добавления расхода.
//GetBalance() float64 для получения текущего баланса.
//12. Валютообменник
//Описание: Реализуйте интерфейс CurrencyConverter с методами ToUSD(amount float64) и ToEUR(amount float64). Реализуйте структуру Exchange, которая будет работать с конвертацией валют и реализует этот интерфейс.
//Методы:
//ToUSD(amount float64) float64 для конвертации в доллары США.
//ToEUR(amount float64) float64 для конвертации в евро.
//13. Банковский счет
//Описание: Реализуйте интерфейс BankAccount с методами Deposit(amount float64) и Withdraw(amount float64). Реализуйте структуру Account, которая будет работать с банковским счетом и реализует этот интерфейс.
//Методы:
//Deposit(amount float64) для депозита средств на счет.
//Withdraw(amount float64) для снятия средств со счета.
//GetBalance() float64 для получения текущего баланса.
//14. Налоговый расчет
//Описание: Реализуйте интерфейс TaxCalculator с методами CalculateIncomeTax(income float64) и CalculateVAT(amount float64). Реализуйте структуру SimpleTaxCalculator, которая будет работать с расчетом налогов и реализует этот интерфейс.
//Методы:
//CalculateIncomeTax(income float64) float64 для расчета налога на доход.
//CalculateVAT(amount float64) float64 для расчета НДС.
//15. Ведение счета-фактуры
//Описание: Реализуйте интерфейс Invoice с методами AddItem(name string, price float64) и Total() float64. Реализуйте структуру SimpleInvoice, которая будет работать со счетами и реализует этот интерфейс.
//Методы:
//AddItem(name string, price float64) для добавления товара в счет.
//Total() float64 для получения общей суммы счета.
//16. Управление портфелем акций
//Описание: Реализуйте интерфейс StockPortfolio с методами Buy(symbol string, quantity int) и Sell(symbol string, quantity int). Реализуйте структуру Portfolio, которая будет работать с управлением акциями и реализует этот интерфейс.
//Методы:
//Buy(symbol string, quantity int) для покупки акций.
//Sell(symbol string, quantity int) для продажи акций.
//GetHoldings() map[string]int для получения текущих владений.
//17. Банкомат
//Описание: Реализуйте интерфейс ATM с методами Deposit(amount float64) и Withdraw(amount float64). Реализуйте структуру ATMSystem, которая будет работать с банкоматом и реализует этот интерфейс.
//Методы:
//Deposit(amount float64) для депозита средств в банкомат.
//Withdraw(amount float64) для снятия средств из банкомата.
//GetBalance() float64 для получения текущего баланса.
//18. Управление зарплатой сотрудников
//Описание: Реализуйте интерфейс Payroll с методами AddEmployee(name string, salary float64) и GetSalary(name string) float64. Реализуйте структуру CompanyPayroll, которая будет работать с зарплатами сотрудников и реализует этот интерфейс.
//Методы:
//AddEmployee(name string, salary float64) для добавления сотрудника и его зарплаты.
//GetSalary(name string) float64 для получения зарплаты сотрудника.
//TotalPayroll() float64 для получения общей суммы зарплат.
//19. Управление библиотекой
//Описание: Реализуйте интерфейс Library с методами AddBook(title, author string) и GetBooks() []string. Реализуйте структуру BookLibrary, которая будет работать с библиотекой книг и реализует этот интерфейс.
//Методы:
//AddBook(title, author string) для добавления книги в библиотеку.
//GetBooks() []string для получения списка книг в библиотеке.
//20. Управление заказами в интернет-магазине
//Описание: Реализуйте интерфейс OrderManager с методами AddOrder(orderID string, amount float64) и GetTotalSales() float64. Реализуйте структуру OnlineStore, которая будет работать с заказами и реализует этот интерфейс.
//Методы:
//AddOrder(orderID string, amount float64) для добавления заказа.
//GetTotalSales() float64 для получения общей суммы продаж.
//21. Управление банковскими транзакциями
//Описание: Реализуйте интерфейс TransactionManager с методами AddTransaction(amount float64) и GetStatement() []string. Реализуйте структуру BankAccount, которая будет работать с банковскими транзакциями и реализует этот интерфейс.
//Методы:
//AddTransaction(amount float64) для добавления транзакции.
//GetStatement() []string для получения выписки по счету.

// 22. Управление каталогом продуктов
// Описание: Реализуйте интерфейс ProductCatalog с методами AddProduct(name string, price float64) и GetProductPrice(name string) float64.
// Реализуйте структуру Catalog, которая будет работать с каталогом продуктов и реализует этот интерфейс.
// Методы:
// AddProduct(name string, price float64) для добавления продукта в каталог.
// GetProductPrice(name string) float64 для получения цены продукта.
// GetProducts() map[string]float64 для получения всех продуктов и их цен.
type ProductCatalog interface {
	AddProduct()
	GetProductPrice()
}
type Catalog struct {
	name  string
	price float64
}

//23. Управление заказами в ресторане
//Описание: Реализуйте интерфейс RestaurantOrderManager с методами AddOrder(orderID string, items []string) и GetOrderDetails(orderID string) []string. Реализуйте структуру Restaurant, которая будет работать с заказами и реализует этот интерфейс.
//Методы:
//AddOrder(orderID string, items []string) для добавления заказа.
//GetOrderDetails(orderID string) []string для получения деталей заказа.
//GetAllOrders() map[string][]string для получения всех заказов.
//24. Расчет ипотечного кредита
//Описание: Реализуйте интерфейс MortgageCalculator с методами CalculateMonthlyPayment(principal, annualRate float64, years int) и CalculateTotalPayment(principal, annualRate float64, years int). Реализуйте структуру Mortgage, которая будет работать с расчетом ипотечного кредита и реализует этот интерфейс.
//Методы:
//CalculateMonthlyPayment(principal, annualRate float64, years int) float64 для расчета ежемесячного платежа.
//CalculateTotalPayment(principal, annualRate float64, years int) float64 для расчета общей суммы выплат.
//25. Управление клиентами банка
//Описание: Реализуйте интерфейс CustomerManager с методами AddCustomer(name string, accountNumber string) и GetCustomerDetails(accountNumber string) (string, string). Реализуйте структуру Bank, которая будет работать с клиентами банка и реализует этот интерфейс.
//Методы:
//AddCustomer(name string, accountNumber string) для добавления клиента.
//GetCustomerDetails(accountNumber string) (string, string) для получения деталей клиента.
//GetAllCustomers() map[string]string для получения всех клиентов.
//26. Учет рабочего времени сотрудников
//Описание: Реализуйте интерфейс TimeTracker с методами ClockIn(employeeID string) и ClockOut(employeeID string). Реализуйте структуру WorkTime, которая будет работать с учетом рабочего времени сотрудников и реализует этот интерфейс.
//Методы:
//ClockIn(employeeID string) для регистрации начала работы.
//ClockOut(employeeID string) для регистрации окончания работы.
//GetWorkingHours(employeeID string) float64 для получения отработанных часов.
//27. Система учета студентов
//Описание: Реализуйте интерфейс StudentManager с методами AddStudent(name string, studentID string) и GetStudentDetails(studentID string) (string, string). Реализуйте структуру University, которая будет работать с учетной записью студентов и реализует этот интерфейс.
//Методы:
//AddStudent(name string, studentID string) для добавления студента.
//GetStudentDetails(studentID string) (string, string) для получения данных студента.
//GetAllStudents() map[string]string для получения списка всех студентов.
//28. Управление расходами предприятия
//Описание: Реализуйте интерфейс ExpenseManager с методами AddExpense(category string, amount float64) и GetTotalExpenses(category string) float64. Реализуйте структуру Enterprise, которая будет работать с расходами предприятия и реализует этот интерфейс.
//Методы:
//AddExpense(category string, amount float64) для добавления расхода.
//GetTotalExpenses(category string) float64 для получения общей суммы расходов по категории.
//GetAllExpenses() map[string]float64 для получения всех расходов.
//29. Управление учебными курсами
//Описание: Реализуйте интерфейс CourseManager с методами AddCourse(courseName string, credits int) и GetCourseCredits(courseName string) int. Реализуйте структуру EducationInstitution, которая будет работать с учебными курсами и реализует этот интерфейс.
//Методы:
//AddCourse(courseName string, credits int) для добавления курса.
//GetCourseCredits(courseName string) int для получения количества кредитов курса.
//GetAllCourses() map[string]int для получения всех курсов.
//30. Учет счетов за коммунальные услуги
//Описание: Реализуйте интерфейс UtilityBillManager с методами AddBill(utilityType string, amount float64) и GetTotalBills(utilityType string) float64.
//Реализуйте структуру Household, которая будет работать с учетом коммунальных услуг и реализует этот интерфейс.
//Методы:
//AddBill(utilityType string, amount float64) для добавления счета за коммунальные услуги.
//GetTotalBills(utilityType string) float64 для получения общей суммы по типу коммунальной услуги.
//GetAllBills() map[string]float64 для получения всех счетов.
