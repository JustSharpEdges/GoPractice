package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	owner   string
	balance float64
}

func NewBankAccount(owner string, initialBalance float64) *BankAccount {
	return &BankAccount{
		owner:   owner,
		balance: initialBalance,
	}
}

func (acc *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	acc.balance += amount
	return nil
}

func (acc *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	if amount > acc.balance {
		return errors.New("недостаточно средств")
	}
	acc.balance -= amount
	return nil
}

func (acc BankAccount) GetBalance() float64 {
	return acc.balance
}

func (acc BankAccount) Display() {
	fmt.Printf("Владелец: %s\n", acc.owner)
	fmt.Printf("Баланс: %.2f\n", acc.balance)
}

func main() {
	account := NewBankAccount("Алексей Иванов", 1000)

	account.Display()

	err := account.Deposit(500)
	if err != nil {
		fmt.Println("Ошибка пополнения:", err)
	} else {
		fmt.Println("После пополнения на 500:")
		account.Display()
	}

	err = account.Withdraw(200)
	if err != nil {
		fmt.Println("Ошибка снятия:", err)
	} else {
		fmt.Println("После снятия 200:")
		account.Display()
	}

	err = account.Withdraw(2000)
	if err != nil {
		fmt.Println("Ошибка снятия:", err)
	}
}
