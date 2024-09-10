package main

import (
	"fmt"
	"strconv"
)

func checkOperands(operand string) (int, string) {
	arabik := [11]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	rim := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, r := range arabik {
		if operand == r {
			intOperand, _ := strconv.Atoi(operand)
			return intOperand, "a"
		}
	}

	for ind, r := range rim {
		if operand == r {
			intOperand := ind + 1
			return intOperand, "r"
		}

	}
	return -1, "error"
}

func getRezultDec(a, b int, oper string) int {
	if oper == "+" {
		return a + b
	}
	if oper == "-" {
		return a - b
	}
	if oper == "*" {
		return a * b
	}
	if oper == "/" {
		if b == 0 {
			panic("деление на 0 запрещено")
		}
		return a / b
	}
	panic("непонятная операция")
}

func getRezult(rez int, a string) string {
	rim := [100]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
		"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
		"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
		"XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
		"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L",
		"LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
		"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
		"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
		"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
		"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
	if a == "r" {
		if rez < 1 {
			panic("в римской системе нет отрицательных чисел")
		}
		rezult := rim[rez-1]
		return rezult
	}
	return strconv.Itoa(rez)
}

func main() {
	var operand1, oper1, operand2, oper2 string
	fmt.Scanf("%s %s %s %s", &operand1, &oper1, &operand2, &oper2)

	if oper2 != "" {
		panic("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
	if oper1 == "" {
		panic("строка не является математической операцией")
	}

	a, typeOperandA := checkOperands(operand1)
	b, typeOperandB := checkOperands(operand2)

	if typeOperandA != typeOperandB {
		panic("используются одновременно разные системы счисления")
	}
	if a == -1 || b == -1 {
		panic("операнды больше 10")
	}
	if a == 0 || b == 0 {
		panic("операнды меньше 1")
	}

	rez := getRezultDec(a, b, oper1)
	rezult := getRezult(rez, typeOperandA)

	fmt.Println(rezult)
}
