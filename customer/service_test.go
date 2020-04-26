package customer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"../account"
	"../customer"
)

func TestAccountService(t *testing.T) {
	mockAccounter := &mockAccounter{accounts: make(map[string]*account.Account), bankBalancer: mockBankBalancer{}}
	customerService := customer.NewService(mockAccounter)
	customerId := customerService.Create()

	t.Run("create customer", func(t *testing.T) {
		assert.NotNil(t, customerId)
	})
	t.Run("deposit - success", func(t *testing.T) {
		err := customerService.Deposit(customerId, 20.0)
		assert.Nil(t, err)
	})
	t.Run("deposit - failure - incorrect customerId", func(t *testing.T) {
		err := customerService.Deposit("gibberish", 20.0)
		assert.NotNil(t, err)
	})
	t.Run("withdraw - success", func(t *testing.T) {
		err := customerService.Withdraw(customerId, 20.0)
		assert.Nil(t, err)
	})
	t.Run("withdraw - failure - incorrect customerId", func(t *testing.T) {
		err := customerService.Withdraw("gibberish", 20.0)
		assert.NotNil(t, err)
	})
	t.Run("balance - success", func(t *testing.T) {
		balance, err := customerService.CurrentBalance(customerId)
		assert.Nil(t, err)
		assert.Equal(t, 0.0, *balance)
	})
	t.Run("balance - failure - incorrect customerId", func(t *testing.T) {
		balance, err := customerService.CurrentBalance("gibberish")
		assert.NotNil(t, err)
		assert.Nil(t, balance)
	})
}

type mockBankBalancer struct {
	balance float64
}

func (m *mockBankBalancer) UpdateBalance(amount float64) {
	m.balance += amount
}

func (m *mockBankBalancer) CurrentBalance() float64 {
	return m.balance
}

type mockAccounter struct {
	accounts     map[string]*account.Account
	bankBalancer mockBankBalancer
}

func (m *mockAccounter) Withdraw(accountId string, customerId string, amount float64) error {
	return nil
}

func (m *mockAccounter) Deposit(accountId string, customerId string, amount float64) error {
	return nil
}

func (m *mockAccounter) Create(customerId string) string {
	return ""
}

func (m *mockAccounter) Balance(accountId string, customerId string) (*float64, error) {
	balance := 0.0
	return &balance, nil
}
