package main

import "fmt"

type BankAccount interface {
	GetBalance() int
	Deposit(amount int)
	WithDraw(amount int) error
}

func main() {

	accounts := []BankAccount{
		NewSBIaccount(),
		NewYesaccount(),
	}

	for _, ac := range accounts {
		ac.Deposit(30)
		if err := ac.WithDraw(30); err != nil {
			panic(err)
		}
		currentBalance := ac.GetBalance()
		fmt.Println("Balance is ", currentBalance)
	}

}
