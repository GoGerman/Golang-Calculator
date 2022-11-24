package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	var primer string
	primer = scan()
	primer = str(primer)

	var sl []string
	for _, i := range primer {
		sl = append(sl, fmt.Sprintf("%c", i))
	}

	var simvol string
	var count int
	simvol, count = getSimvol(sl) // получение знака операции
	if count == 0 {               // проверка
		fmt.Println("Вывод ошибки, так как строка не является математической операцией")
		runtime.Goexit()
	} else if count >= 2 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		runtime.Goexit()
	}

	var part1, part2 string
	part1 = p1(part1, sl, simvol)
	part2 = p2(part2, sl, simvol)

	var num1, num2 int
	num1, err1 := strconv.Atoi(part1)
	num2, err2 := strconv.Atoi(part2)

	if (part1 == "" || part2 == "") || (part1 == "" && part2 == "") {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией")
		runtime.Goexit()
	}

	if err1 != nil || err2 != nil { // римские
		var rom1, rom2 int
		var romB1, romB2 bool
		rom1, romB1 = RomanToInt(part1)
		rom2, romB2 = RomanToInt(part2)
		if romB1 == false || romB2 == false {
			fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления или строка не является математической операцией")
			fmt.Println("Так же, возможно< вы используете числа больше 10 (Х)")
			runtime.Goexit()
		}
		var result1 = sim(simvol, rom1, rom2)

		if result1 <= 0 {
			fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел")
			runtime.Goexit()
		}
		result2 := IntToRoman(result1)
		fmt.Println(result2)
	} else { // арабские
		if test1(num1) == true || test1(num2) == true {
			fmt.Println("Вывод ошибки, так как калькулятор не принимает числа меньше 1 или больше 10")
			runtime.Goexit()
		}
		var result3 = sim(simvol, num1, num2)
		fmt.Println(result3)
	}
}

func scan() string { // сканирует строку
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()
}

func str(a string) string { // убирает пробелы
	a = strings.TrimSpace(a)
	a = strings.ReplaceAll(a, " ", "")
	return a
}

func p1(a string, b []string, c string) string { // 1 число
	for _, j := range b {
		if j == c {
			break
		}
		a += j
	}
	return a
}

func p2(a string, b []string, c string) string { // 2 число и реверс его
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == c {
			break
		}
		a += b[i]
	}
	var a1 string
	for _, j := range a {
		a1 = string(j) + a1
	}
	return a1
}

func getSimvol(sl []string) (string, int) { // получение символа для операции
	var simvol string
	count := 0
	for _, i := range sl {
		if i == "+" || i == "-" || i == "*" || i == "/" {
			simvol = i
			count += 1
		}
	}
	return simvol, count
}

func test1(num int) bool { // тест на больше 10
	if num > 10 || num < 1 {
		return true
	}
	return false
}

func RomanToInt(str string) (int, bool) { // перевод римских в арабские
	conversions := []struct {
		digit string
		value int
	}{
		{"X", 10},
		{"IX", 9},
		{"VIII", 8},
		{"VII", 7},
		{"VI", 6},
		{"V", 5},
		{"IV", 4},
		{"III", 3},
		{"II", 2},
		{"I", 1},
	}
	nol := 0
	for _, conversion := range conversions {
		if conversion.digit == str {
			return conversion.value, true
		}
	}
	return nol, false
}

func IntToRoman(number int) string { // перевод арабских в римские
	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

func sim(simvol string, num1, num2 int) int {
	var result int
	switch simvol {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}
	return result
}
