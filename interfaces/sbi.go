package main

import "errors"

type SbiAccount struct {
	Balance int
}

func NewSBIaccount() *SbiAccount {
	return &SbiAccount{
		Balance: 0,
	}
}

func (s *SbiAccount) GetBalance() int {
	return s.Balance
}

func (s *SbiAccount) WithDraw(amount int) error {
	newbalance := s.Balance - amount
	if newbalance >= s.Balance {
		return errors.New("balanace is less")
	}
	s.Balance = newbalance
	return nil

}

func (s *SbiAccount) Deposit(amount int) {
	s.Balance += amount
}
