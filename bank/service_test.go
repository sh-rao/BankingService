package bank_test

import (
	"testing"

	"../bank"

	"github.com/stretchr/testify/assert"
)

func TestBankService(t *testing.T) {
	bankService := bank.NewService(0.0)
	t.Run("update balance - check balance - postive amount", func(t *testing.T) {
		bankService.UpdateBalance(10.0)
		assert.Equal(t, 10.00, bankService.CurrentBalance())
	})
	t.Run("update balance - check balance - negative amount", func(t *testing.T) {
		bankService.UpdateBalance(-5.0)
		assert.Equal(t, 5.00, bankService.CurrentBalance())
	})
}
