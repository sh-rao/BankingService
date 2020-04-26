package account

import (
	"errors"
	"math"

	"github.com/google/uuid"
)

type Service struct {
	accounts     map[string]*Account
	bankBalancer BankBalancer
}

type BankBalancer interface {
	UpdateBalance(amount float64)
}

func NewService(bankBalancer BankBalancer) *Service {
	return &Service{bankBalancer: bankBalancer, accounts: make(map[string]*Account)}
}

func (s *Service) Create(customerId string) string {
	accountId := uuid.New().String()
	s.accounts[accountId] = &Account{id: accountId, customerId: customerId}
	return accountId
}

func (s *Service) Withdraw(accountId string, customerId string, amount float64) error {
	if s.accounts[accountId] == nil {
		return errors.New("could not find an account with the given account id")
	}
	account := s.accounts[accountId]
	if account.customerId != customerId {
		return errors.New("account and customer ids do not match")
	}
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than 0")
	}
	if account.balance < amount {
		return errors.New("account balance is lesser than the withdrawal amount")
	}
	account.balance -= amount
	account.balance = math.Round(account.balance*100) / 100
	s.bankBalancer.UpdateBalance(-amount)
	return nil
}

func (s *Service) Deposit(accountId string, customerId string, amount float64) error {
	if s.accounts[accountId] == nil {
		return errors.New("could not find an account with the given account id")
	}
	account := s.accounts[accountId]
	if account.customerId != customerId {
		return errors.New("account and customer ids do not match")
	}
	if amount <= 0 {
		return errors.New("deposit amount must be greater than 0")
	}
	account.balance += amount
	account.balance = math.Round(account.balance*100) / 100
	s.bankBalancer.UpdateBalance(amount)
	return nil
}

func (s *Service) Balance(accountId string, customerId string) (*float64, error) {
	if s.accounts[accountId] == nil {
		return nil, errors.New("could not find an account with the given account id")
	}
	account := s.accounts[accountId]
	if account.customerId != customerId {
		return nil, errors.New("account and customer ids do not match")
	}
	return &account.balance, nil
}
