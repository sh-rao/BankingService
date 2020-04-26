package customer

import (
	"errors"

	"github.com/google/uuid"
)

type Service struct {
	customers map[string]*Customer
	accounter Accounter
}

type Accounter interface {
	Balance(accountId string, customerId string) (*float64, error)
	Withdraw(accountId string, customerId string, amount float64) error
	Deposit(accountId string, customerId string, amount float64) error
	Create(customerId string) string
}

func NewService(accounter Accounter) *Service {
	return &Service{accounter: accounter, customers: make(map[string]*Customer)}
}

func (s *Service) Create() string {
	customerId := uuid.New().String()
	accountId := s.accounter.Create(customerId)
	s.customers[customerId] = &Customer{id: customerId, accountId: accountId}
	return customerId
}

func (s *Service) CurrentBalance(customerId string) (*float64, error) {
	if s.customers[customerId] == nil {
		return nil, errors.New("could not find a customer with the given id")
	}
	return s.accounter.Balance(s.customers[customerId].accountId, customerId)
}

func (s *Service) Withdraw(customerId string, amount float64) error {
	if s.customers[customerId] == nil {
		return errors.New("could not find a customer with the given id")
	}
	return s.accounter.Withdraw(s.customers[customerId].accountId, customerId, amount)
}

func (s *Service) Deposit(customerId string, amount float64) error {
	if s.customers[customerId] == nil {
		return errors.New("could not find a customer with the given id")
	}
	return s.accounter.Deposit(s.customers[customerId].accountId, customerId, amount)
}
