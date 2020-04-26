package bank

import "math"

type Service struct {
	balance float64
}

func NewService(initialBalance float64) *Service {
	return &Service{balance: initialBalance}
}

func (s *Service) UpdateBalance(amount float64) {
	s.balance += amount
	s.balance = math.Round(s.balance*100) / 100
}

func (s *Service) CurrentBalance() float64 {
	return s.balance
}
