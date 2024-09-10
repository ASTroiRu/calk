package main

import (
	"fmt"
	"strconv"
)

func checkOperands(operand string) (int, string) {
	arifm := [11]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	rim := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, r := range arifm {
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
	rim := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
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
	rez := getRezultDec(a, b, oper1)
	rezult := getRezult(rez, typeOperandA)

	fmt.Println(rezult)
}
