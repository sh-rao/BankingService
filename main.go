package main

import (
	"fmt"

	"./account"
	"./bank"
	"./customer"
)

func main() {
	bankService := bank.NewService(0)
	accountService := account.NewService(bankService)
	customerService := customer.NewService(accountService)

	var customerIds [5]string
	for i := 0; i < 5; i++ {
		customerIds[i] = customerService.Create()
	}

	customerService.Deposit(customerIds[0], 10.00)
	customerService.Deposit(customerIds[1], 20.00)
	customerService.Deposit(customerIds[2], 30.00)
	//This will fail as it's a negative amount
	customerService.Deposit(customerIds[3], -10.00)
	customerService.Deposit(customerIds[4], 20.00)

	customerService.Withdraw(customerIds[0], 5.00)
	customerService.Withdraw(customerIds[1], 15.00)
	customerService.Withdraw(customerIds[2], 25.00)
	//These two will fail due to insufficient balance
	customerService.Withdraw(customerIds[3], 10.00)
	customerService.Withdraw(customerIds[4], 21.00)

	for i := 0; i < 5; i++ {
		balance, err := customerService.CurrentBalance(customerIds[i])
		if err != nil {
			fmt.Println("error: ", err)
		}
		fmt.Printf("Customer %s Balance: %.2f\n", customerIds[i], *balance)
	}

	fmt.Printf("Bank's Balance: %.2f\n", bankService.CurrentBalance())
}
