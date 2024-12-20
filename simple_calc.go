/* Проект выполнен в рамках учебного курса.
Представляет собой простой калькулятор, который может:
принимать два числа с плавающей точкой
принимать символ, который обозначает операцию, которую пользователь хочет совершить над числами
вычислять результат операции
выводит результат на консоль

Операции, которые калькулятор выполняет:
Сложение +
Вычитание -
Умножение *
Деление /
Остаток от деления %
*/

package main 

import (
	"fmt" 
)

func main() { 
	var num1, num2 float64 
	var operation string   

	_, err := fmt.Scan(&num1)            // сканируем первое число
	if err != nil {                      // если введено не число, то выводим предупреждение
		fmt.Println("Ошибка ввода! Убедитесь, что вы вводите числа.")
		return
	}

	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Ошибка ввода! Убедитесь, что вы вводите числа.")
		return
	}

	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Ошибка ввода! Убедитесь, что вы вводите символ операции.")
		return
	}

	switch operation {
	case "+":
		fmt.Print(num1+num2)
	case "-":
		fmt.Print(num1-num2)
	case "*":
		fmt.Print(num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Print(num1/num2)
		}
	case "%":
		// Приводим к int для операции остатка от деления
		intNum1 := int(num1)
		intNum2 := int(num2)
		if intNum2 == 0 {
			fmt.Println("Делить на ноль нельзя!")
		} else {
			fmt.Print(intNum1%intNum2)
		}
	default:
		fmt.Println("Неизвестная операция. Попробуйте снова.")
	}
}
