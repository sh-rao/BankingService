package account_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"../account"
)

func TestAccountService(t *testing.T) {
	accountService := account.NewService(&mockBankBalancer{balance: 0})
	customerId := uuid.New().String()
	accountId := accountService.Create(customerId)
	t.Run("create account", func(t *testing.T) {
		assert.NotNil(t, accountId)
	})
	t.Run("deposit account & balance test - success - postive amount", func(t *testing.T) {
		err := accountService.Deposit(accountId, customerId, 20.0)
		assert.Nil(t, err)
		balance, _ := accountService.Balance(accountId, customerId)
		assert.Equal(t, *balance, 20.0)
	})
	t.Run("deposit account - failure - bad customerid", func(t *testing.T) {
		err := accountService.Deposit(accountId, "gibberish", 10.0)
		assert.NotNil(t, err)
	})
	t.Run("deposit account - failure - bad accountId", func(t *testing.T) {
		err := accountService.Deposit("gibberish", customerId, 10.0)
		assert.NotNil(t, err)
	})
	t.Run("deposit account & balance test - failure - amount equals 0", func(t *testing.T) {
		err := accountService.Deposit(accountId, customerId, 0.0)
		assert.NotNil(t, err)
		balance, _ := accountService.Balance(accountId, customerId)
		assert.Equal(t, *balance, 20.0)
	})
	t.Run("deposit account & balance - failure - amount less than 0", func(t *testing.T) {
		err := accountService.Deposit(accountId, customerId, -10.0)
		assert.NotNil(t, err)
		balance, _ := accountService.Balance(accountId, customerId)
		assert.Equal(t, *balance, 20.0)
	})
	t.Run("deposit account & balance - success - positive decimal amount", func(t *testing.T) {
		err := accountService.Deposit(accountId, customerId, 10.10)
		assert.Nil(t, err)
		balance, _ := accountService.Balance(accountId, customerId)
		assert.Equal(t, *balance, 30.10)
	})
	t.Run("withdraw account & balance test - success - postive amount", func(t *testing.T) {
		err := accountService.Withdraw(accountId, customerId, 10.0)
		assert.Nil(t, err)
		balance, _ := accountService.Balance(accountId, customerId)
		assert.Equal(t, *balance, 20.10)
	})
	t.Run("withdraw account - failure - bad customerid", func(t *testing.T) {
		err := accountService.Withdraw(accountId, "gibberish", 10.0)
		assert.NotNil(t, err)
	})
	t.Run("withdraw account - failure - bad accountId", func(t *testing.T) {
		err := accountService.Withdraw("gibberish", customerId, 10.0)
		assert.NotNil(t, err)
	})
	t.Run("withdraw account & balance test - failure - amount equals 0", func(t *testing.T) {
		err := accountService.Withdraw(accountId, customerId, 0.0)
		assert.NotNil(t, err)
		balance, _ := accountService.Balance(accountId, customerId)
		assert.Equal(t, *balance, 20.10)
	})
	t.Run("withdraw account & balance - failure - amount less than 0", func(t *testing.T) {
		err := accountService.Withdraw(accountId, customerId, -10.0)
		assert.NotNil(t, err)
		balance, _ := accountService.Balance(accountId, customerId)
		assert.Equal(t, *balance, 20.10)
	})
	t.Run("withdraw account & balance - success - positive decimal amount", func(t *testing.T) {
		err := accountService.Withdraw(accountId, customerId, 10.10)
		assert.Nil(t, err)
		balance, _ := accountService.Balance(accountId, customerId)
		assert.Equal(t, *balance, 10.0)
	})
	t.Run("balance - failure - bad customerid", func(t *testing.T) {
		balance, err := accountService.Balance(accountId, "gibberish")
		assert.Nil(t, balance)
		assert.NotNil(t, err)
	})
	t.Run("balance - failure - bad accountId", func(t *testing.T) {
		balance, err := accountService.Balance("gibberish", customerId)
		assert.Nil(t, balance)
		assert.NotNil(t, err)
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
