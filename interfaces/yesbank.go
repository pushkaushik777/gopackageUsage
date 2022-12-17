package main

import "errors"

type YesAccount struct {
	Balance int
	Fee     int
}

func NewYesaccount() *YesAccount {
	return &YesAccount{
		Balance: 0,
		Fee:     40,
	}
}

func (s *YesAccount) GetBalance() int {
	return s.Balance
}

func (s *YesAccount) WithDraw(amount int) error {
	newbalance := s.Balance - amount - s.Fee
	if newbalance >= s.Balance {
		return errors.New("balanace is less")
	}
	s.Balance = newbalance
	return nil

}

func (s *YesAccount) Deposit(amount int) {
	s.Balance += amount
}
