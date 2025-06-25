package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (acc *BankAccount) deposit(amount float64) {
	if amount > 0 {
		acc.Balance += amount
		fmt.Printf("Пополнение на %.2f. Баланс изменен: %.2f\n", amount, acc.Balance)
	} else {
		fmt.Println("Ошибка: сумма должна быть положительной")
	}
}

func (acc *BankAccount) withdraw(amount float64) {
	if amount > 0 && amount <= acc.Balance {
		acc.Balance -= amount
		fmt.Printf("Было снято %.2f. Остаток: %.2f\n", amount, acc.Balance)
	} else {
		fmt.Println("Ошибка: недостаточно средств или неверная сумма")
	}
}

func (acc BankAccount) info() {
	fmt.Printf("Пользователь %s имеет баланс = %.2f\n", acc.Owner, acc.Balance)
}

func main() {
	account := BankAccount{Owner: "Елисеенко Владислав", Balance: 5678.13}

	account.info()
	account.deposit(560)
	account.withdraw(200)
	account.withdraw(10000)
}
